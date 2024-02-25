{
  containerDefinitions: [
    {
      cpu: 0,
      essential: true,
      // image: 'nginx',
      image: 'ghcr.io/magi-sche-org/backend/server:build-20240223-62edc4b2',
      name: 'magische-dev-api',
      portMappings: [
        {
          appProtocol: '',
          containerPort: 80,
          hostPort: 80,
          protocol: 'tcp',
        },
      ],
    },
  ],
  cpu: '256',
  executionRoleArn: 'arn:aws:iam::905418376731:role/magische-dev-api-server-task-exec',
  family: 'magische-dev-api',
  ipcMode: '',
  memory: '512',
  networkMode: 'awsvpc',
  pidMode: '',
  requiresCompatibilities: [
    'FARGATE',
  ],
  runtimePlatform: {
    cpuArchitecture: 'ARM64',
    operatingSystemFamily: 'LINUX',
  },
  tags: [
    {
      key: 'Env',
      value: 'dev',
    },
    {
      key: 'Service',
      value: 'api',
    },
    {
      key: 'Name',
      value: 'magische-dev-api-server',
    },
  ],
  taskRoleArn: 'arn:aws:iam::905418376731:role/magische-dev-api-server-task',
}
