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
    req = ListRequest.new({title: ""})
    @stub.list_tasks(req).tasks
  end
end
