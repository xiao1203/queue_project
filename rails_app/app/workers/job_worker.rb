require 'redis'

class JobWorker
  include Sidekiq::Worker

  def perform
    Rails.logger.debug "JobWorker started"
    
    redis_url = ENV['REDIS_URL'] || 'redis://redis:6379'
    redis = Redis.new(url: redis_url)

    job_json = redis.lpop('job_queue')
    if job_json
      job = JSON.parse(job_json)
      process_job(job)
    end
  end

  private

  def process_job(job)
    puts "Processing job: #{job['id']}"
    # 実際の処理をここに記述
  end
end
