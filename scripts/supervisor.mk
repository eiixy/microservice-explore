.PHONY: supervisor.%

# 服务启动端口
SERVER_PORTS := 8081 8082 8083 8084

# 安装 supervisor
supervisor.install:
	apt-get update
	apt-get install supervisor

# 查看状态
supervisor.status:
	supervisorctl status

supervisor.start:
	supervisorctl start short-link-job
	for port in $(SERVER_PORTS);do supervisorctl start short-link-$$port; done

supervisor.stop:
	supervisorctl stop short-link-job
	for port in $(SERVER_PORTS);do supervisorctl stop short-link-$$port; done

# 重启
supervisor.restart:
	supervisorctl restart short-link-job
	for port in $(SERVER_PORTS);do supervisorctl restart short-link-$$port; done

supervisor.reload:
	supervisorctl reload

supervisor.run:
	supervisord -c /etc/supervisor/supervisord.conf
