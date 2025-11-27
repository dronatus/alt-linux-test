.PHONY: apply delete describe status test run stop-run status-run clean

PORT_FILE := .current_port
PID_FILE := .portforward_pid
STATUS_FILE := .status

apply:
	kubectl apply -f k8s/
	@echo 1 > $(STATUS_FILE)

delete:
	kubectl delete -f k8s/
	@rm -f $(STATUS_FILE)

describe:
	kubectl describe deployment test-quiz

status:
	kubectl get pods,services,deployments

test:
	@if [ ! -f $(STATUS_FILE) ] || [ $(STATUS_FILE) -ne 1 ]; then \
		echo "Can't run: Deployment is not apply."; \
		echo "Run 'make apply' first."; \
		exit 1; \
	fi
	@echo "Testing application via port-forward..."
	@rm -f $(PORT_FILE)
	@for port in 8080 8000 8081 8082; do \
		echo "Trying port $$port..."; \
		kubectl port-forward service/test-service $$port:61111 & \
		PORT_PID=$$!; \
		sleep 2; \
		if curl -s http://localhost:$$port > /dev/null; then \
			echo "Successfully connected on port $$port"; \
			echo "=== Application Output ==="; \
			curl -s http://localhost:$$port | head -n 5; \
			echo "=========================="; \
			echo $$port > $(PORT_FILE); \
			kill $$PORT_PID 2>/dev/null || true; \
			break; \
		else \
			echo "Port $$port failed, trying next..."; \
			kill $$PORT_PID 2>/dev/null || true; \
		fi; \
	done
	@if [ ! -f $(PORT_FILE) ]; then \
		echo "ERROR: All ports failed!"; \
		exit 1; \
	fi
	@echo "Test completed!"
	@echo ""

run: test
	@if [ ! -f $(PORT_FILE) ]; then \
		echo "Can't run: No port found."; \
		exit 1; \
	fi
	@PORT=$$(cat $(PORT_FILE)); \
	echo ""; \
	echo "Starting background port-forward on port $$PORT..."; \
	kubectl port-forward service/test-service $$PORT:61111 & \
	echo $$! > $(PID_FILE); \
	echo "Port-forward running in background (PID: $$!)"; \
	echo "Access your application at: http://localhost:$$PORT"; \
	echo "To stop, run: make stop-run"

stop-run:
	@if [ -f $(PID_FILE) ]; then \
		PID=$$(cat $(PID_FILE)); \
		if ps -p $$PID > /dev/null; then \
			kill $$PID; \
			echo "Stopped port-forward (PID: $$PID)"; \
		else \
			echo "Process $$PID not running"; \
		fi; \
		rm -f $(PID_FILE); \
	else \
		echo "No running port-forward found"; \
	fi
	@rm -f $(PORT_FILE)

status-run:
	@if [ -f $(PID_FILE) ] && [ -f $(PORT_FILE) ]; then \
		PID=$$(cat $(PID_FILE)); \
		PORT=$$(cat $(PORT_FILE)); \
		if ps -p $$PID > /dev/null; then \
			echo "Port-forward running (PID: $$PID, PORT: $$PORT)"; \
			echo "Access: http://localhost:$$PORT"; \
		else \
			echo "Port-forward PID exists but process not running"; \
		fi; \
	else \
		echo "No port-forward running"; \
	fi

clean:
	@$(MAKE) stop-run
	@rm -f $(PORT_FILE) $(PID_FILE)
	@echo "Cleaned up all temporary files"