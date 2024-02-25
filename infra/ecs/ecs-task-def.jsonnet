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
