// This is a template for PM2 ecosystem.config.js

module.exports = {
  apps: [
    {
      name: 'rag-service',
      script: '/usr/bin/go',
      args: 'run main.go',
      exec_interpreter: 'none',
      exec_mode: 'fork',
      watch: true,
      time: true,
      env: {}
    }
  ]
}
