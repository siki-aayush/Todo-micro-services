import { useState, useEffect } from 'react';
import './App.css';

interface Todo {
  id: string;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}
const BASE_URL = 'http://localhost:8081';
function App() {
  const [task, setTask] = useState('');
  const [description, setDescription] = useState('');
  const [tasks, setTasks] = useState<Todo[]>([]);
  const [isEditing, setIsEditing] = useState<string | null>(null);
  const [editTask, setEditTask] = useState('');

  // Fetch todos from the backend
  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const response = await fetch(`${BASE_URL}/todo`);
        if (response.ok) {
          const data = await response.json();
          if (!data.error) {
            setTasks(data.data); // Set the tasks from the API response
          } else {
            console.error('Error fetching todos:', data.message);
          }
        } else {
          console.error('Failed to fetch todos');
        }
      } catch (error) {
        console.error('Error fetching todos:', error);
      }
    };

    fetchTodos();
  }, []);

  const addTask = async () => {
    if (task.trim() && description.trim()) {
      const newTask = {
        name: task,
        description: description,
      };
  
      try {
        const response = await fetch(`${BASE_URL}/todo`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(newTask),
        });
  
        if (response.ok) {
          const data = await response.json();
          console.log('Full API Response:', data); // Log the full response
          if (!data.error && data.data) {
            setTasks((prevTasks) => [data.data, ...prevTasks]); // Prepend the new task to the list
            setTask(''); // Clear the task input
            setDescription(''); // Clear the description input
          } else {
            console.error('Error adding todo:', data.message || 'Unexpected response format');
          }
        } else {
          console.error('Failed to add todo');
        }
      } catch (error) {
        console.error('Error adding todo:', error);
      }
    }
  };

  const deleteTask = async (id: string) => {
    try {
      const response = await fetch(`${BASE_URL}/todo`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id }),
      });
  
      if (response.ok) {
        const data = await response.json();
        console.log('Delete API Response:', data); // Log the API response
        if (!data.error) {
          setTasks((prevTasks) => prevTasks.filter((task) => task.id !== id)); // Remove the task from the list
        } else {
          console.error('Error deleting todo:', data.message);
        }
      } else {
        console.error('Failed to delete todo');
      }
    } catch (error) {
      console.error('Error deleting todo:', error);
    }
  };

  const updateTask = () => {
    if (editTask.trim() && isEditing !== null) {
      const updatedTasks = tasks.map((task) =>
        task.id === isEditing ? { ...task, name: editTask, updated_at: new Date().toISOString() } : task
      );
      setTasks(updatedTasks);
      setIsEditing(null);
      setEditTask('');
    }
  };

  return (
    <div className="App">
      <h1>Todo Application</h1>
      <div className="input-container">
        <input
          type="text"
          placeholder="Enter a task"
          value={task}
          onChange={(e) => setTask(e.target.value)}
        />
        <input
          type="text"
          placeholder="Enter a description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button onClick={isEditing !== null ? updateTask : addTask}>
          {isEditing !== null ? "Update Task" : "Add Task"}
        </button>
      </div>
        <ul className="task-list">
        {tasks.map((task) => (
            <li key={task.id} className="task-item">
            <span>{task.name}</span>
            <p>{task.description}</p>
            <div className="task-actions">
                <button onClick={() => deleteTask(task.id)}>Delete</button>
            </div>
            </li>
        ))}
        </ul>
    </div>
  );
}

export default App;