require "todo_services"

class TodoClient
  def initialize(stub)
    @stub = stub
  end

  def add(title, description)
    task = Task.new(title: title, description: description)
    @stub.add_task(task)
  end

  def list
    @stub.list_tasks(Void.new).tasks
  end
end
