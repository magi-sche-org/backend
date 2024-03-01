{
  containerDefinitions: [
    {
      cpu: 0,
      essential: true,
      image: '{{ must_env `IMAGE_NAME` }}',
      logConfiguration: {
        logDriver: 'awslogs',
        options: {
          'awslogs-group': '/ecs/magische-{{ must_env `ENV` }}-api-server',
          'awslogs-region': '{{ must_env `AWS_REGION` }}',
          'awslogs-stream-prefix': 'magische-{{ must_env `ENV` }}-api-server',
        },
      },
      name: 'magische-{{ must_env `ENV` }}-api',
      portMappings: [
        {
          appProtocol: '',
          containerPort: 8080,
          hostPort: 8080,
          protocol: 'tcp',
        },
      ],
      environment: [
        {
          name: 'ENV',
          value: '{{ must_env `ENV` }}',
        },
        {
          name: 'BASE_URL',
          value: '{{ must_env `BASE_URL` }}',
        },
        {
          name: 'PORT',
          value: '{{ must_env `PORT` }}',
        },
        {
          name: 'SQL_LOG',
          value: '{{ must_env `SQL_LOG` }}',
        }
        {
          name: 'ACCESS_TOKEN_EXPIRE_MINUTES',
          value: '{{ must_env `ACCESS_TOKEN_EXPIRE_MINUTES` }}',
        },
        {
          name: 'REFRESH_TOKEN_EXPIRE_MINUTES',
          value: '{{ must_env `REFRESH_TOKEN_EXPIRE_MINUTES` }}',
        },
        {
          name: 'MYSQL_HOST',
          value: '{{ tfstate `output.rds_endpoint` }}',
        },
        {
          name: 'MYSQL_PORT',
          value: '{{ tfstate `output.rds_port` }}',
        },
        {
          name: 'MYSQL_DATABASE',
          value: '{{ tfstate `output.rds_db_name` }}',
        },
        {
          name: 'CSRF_DISABLED',
          value: 'true',
        },
        {
          name: 'OAUTH_DEFAULT_RETURN_URL',
          value: '{{ must_env `OAUTH_DEFAULT_RETURN_URL` }}',
        },
        {
          name: 'CORS_ORIGINS',
          value: '{{ must_env `CORS_ORIGINS` }}',
        },
      ],
      secrets: [
        {
          name: 'MYSQL_USER',
          valueFrom: '{{ tfstate `output.rds_admin_password_secret_arn` }}:username',
        },
        {
          name: 'MYSQL_PASSWORD',
          valueFrom: '{{ tfstate `output.rds_admin_password_secret_arn` }}:password',
        },
        {
          name: 'SECRET_KEY',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:secret_key',
        },
        {
          name: 'OAUTH_GOOGLE_CLIENT_ID',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:oauth_google_client_id',
        },
        {
          name: 'OAUTH_GOOGLE_CLIENT_SECRET',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:oauth_google_client_secret',
        },
        {
          name: 'OAUTH_MICROSOFT_CLIENT_ID',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:oauth_microsoft_client_id',
        },
        {
          name: 'OAUTH_MICROSOFT_CLIENT_SECRET',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:oauth_microsoft_client_secret',
        },
        {
          name: 'SMTP_EMAIL',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:smtp_email',
        },
        {
          name: 'SMTP_USER',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:smtp_user',
        }
        {
          name: 'SMTP_PASSWORD',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:smtp_password',
        },
        {
          name: 'SMTP_HOST',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:smtp_host',
        },
        {
          name: 'SMTP_PORT',
          valueFrom: '{{ tfstate `output.api_server_ssm_arn` }}:smtp_port',
        },
      ],
    },
  ],
  cpu: '{{ must_env `CPU` }}',
  executionRoleArn: 'arn:aws:iam::905418376731:role/magische-{{ must_env `ENV` }}-api-server-task-exec',
  family: 'magische-{{ must_env `ENV` }}-api',
  ipcMode: '',
  memory: '{{ must_env `MEMORY` }}',
  networkMode: 'awsvpc',
  pidMode: '',
  requiresCompatibilities: [
    'FARGATE',
  ],
  runtimePlatform: {
    cpuArchitecture: '{{ must_env `CPU_ARCHITECTURE` }}',
    operatingSystemFamily: 'LINUX',
  },
  tags: [
    {
      key: 'Env',
      value: '{{ must_env `ENV` }}',
    },
    {
      key: 'Service',
      value: 'api',
    },
    {
      key: 'Name',
      value: 'magische-{{ must_env `ENV` }}-api-server',
    },
  ],
  taskRoleArn: 'arn:aws:iam::905418376731:role/magische-{{ must_env `ENV` }}-api-server-task',
}
