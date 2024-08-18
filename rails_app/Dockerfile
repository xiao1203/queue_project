FROM ruby:3.0.7

RUN gem install bundler:2.3.10

WORKDIR /app

COPY Gemfile ./
RUN bundle install

COPY . .

CMD ["rails", "server", "-b", "0.0.0.0"]
