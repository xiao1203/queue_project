Sidekiq.configure_server do |config|
  config.on(:startup) do
    Sidekiq::Cron::Job.create(
      name: 'Job Queue Check',
      cron: '* * * * *', # 毎分実行
      class: 'JobWorker'
    )
  end
end
