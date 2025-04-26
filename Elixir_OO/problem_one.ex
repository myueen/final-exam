defmodule ProblemOne do
  def new do
    pid_animal_1 = spawn_link(__MODULE__, :animal, ["Rex", "dog"])
    pid_animal_2 = spawn_link(__MODULE__, :animal, ["Whiskers", "cat"])
    send(pid_animal_1, {:speak, pid_animal_1})
    send(pid_animal_2, {:speak, pid_animal_2})

    Process.sleep(1000)
  end

  def animal(name, type) do
    receive do
      {:speak, _pid} ->
        case type do
          "dog" ->
            dog_pid = spawn_link(__MODULE__, :dog, [name])
            send(dog_pid, :speak)
          "cat" ->
            cat_pid = spawn_link(__MODULE__, :cat, [name])
            send(cat_pid, :speak)
        end
    end
  end

  def dog(name) do
    receive do
      :speak ->
        IO.puts("#{name} says: Woof!")
    end
  end

  def cat(name) do
    receive do
      :speak ->
        IO.puts("#{name} says: Meow!")
    end
  end
end
