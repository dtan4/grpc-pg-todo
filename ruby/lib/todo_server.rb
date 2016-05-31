class TodoServer < Todo::Service
  def initialize(conn)
    @conn = conn
  end

  def add_task(task)
    conn.exec("INSERT INTO todos(title, description) VALUES($1, $2)", [task.title, task.description])
    task
  end

  def list_tasks(list_request)
    conn.query("SELECT title, description FROM todos").inject([]) do |todos, column|
      todos << Todo.new(title: column["title"], description: column["description"])
      todos
    end
  end
end
