require "logger"
require "todo_services"

module RubyLogger
  def logger
    LOGGER
  end

  LOGGER = Logger.new(STDOUT)
end

module GRPC
  extend RubyLogger
end

class TodoServer < Todo::Service
  def initialize(conn)
    @conn = conn
  end

  def add_task(task, _call)
    @conn.exec("INSERT INTO todos(title, description) VALUES($1, $2)", [task.title, task.description])
    task
  end

  def list_tasks(list_request, _call)
    tasks = @conn.query("SELECT title, description FROM todos").inject([]) do |result, column|
      result << Task.new(title: column["title"], description: column["description"])
      result
    end

    Tasks.new(tasks: tasks)
  end
end
