{
  capacityProviderStrategy: [
    {
      base: 0,
      capacityProvider: 'FARGATE_SPOT',
      weight: 1,
    },
    {
      base: 1,
      capacityProvider: 'FARGATE',
      weight: 0,
    },
  ],
  deploymentConfiguration: {
    deploymentCircuitBreaker: {
      enable: true,
      rollback: true,
    },
    maximumPercent: 200,
    minimumHealthyPercent: 100,
  },
  deploymentController: {
    type: 'ECS',
  },
  desiredCount: 1,
  enableECSManagedTags: false,
  enableExecuteCommand: false,
  healthCheckGracePeriodSeconds: 0,
  launchType: '',
  loadBalancers: [
    {
      containerName: 'magische-dev-api',
      containerPort: 80,
      targetGroupArn: 'arn:aws:elasticloadbalancing:ap-northeast-1:905418376731:targetgroup/magische-dev-api-4d0baf/a1785e416506a640',
    },
  ],
  networkConfiguration: {
    awsvpcConfiguration: {
      assignPublicIp: 'DISABLED',
      securityGroups: [
        'sg-0d965c0908d641770',
      ],
      subnets: [
        'subnet-0986ad61091b1f3b2',
        'subnet-08da485f0b2da6e19',
        'subnet-01410edfe83dd3cbb',
      ],
    },
  },
  platformFamily: 'Linux',
  platformVersion: 'LATEST',
  propagateTags: 'NONE',
  schedulingStrategy: 'REPLICA',
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
      value: 'magische-dev-api-ecs-service',
    },
  ],
}
