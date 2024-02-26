{
  region: 'ap-northeast-1',
  cluster: 'magische-{{ must_env `ENV` }}',
  service: 'magische-{{ must_env `ENV` }}-api',
  service_definition: '',
  task_definition: 'ecs-task-def.jsonnet',
  timeout: '10m0s',
  // plugins: [
  //   {
  //     name: 'tfstate',
  //     config: {
  //       url: 'remote://app.terraform.io/magische/magische_infra_dev',
  //     },
  //   },
  // ],
}
