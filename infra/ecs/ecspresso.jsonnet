{
  region: 'ap-northeast-1',
  cluster: 'magische-{{ must_env `ENV` }}',
  service: 'magische-{{ must_env `ENV` }}-api',
  // service_definition: 'test/ecs-service-def.jsonnet',
  service_definition: '',
  task_definition: 'ecs-task-def.jsonnet',
  timeout: '10m0s',
  plugins: [
    {
      name: 'tfstate',
      config: {
        url: 'remote://app.terraform.io/magische/{{ must_env `TFC_WORKSPACE` }}',  // like magische_infra_dev
      },
    },
  ],
}
