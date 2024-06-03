import { Box, List, ThemeIcon } from "@mantine/core";
import { Button } from '@material-ui/core';
import { CheckCircle, RadioButtonUnchecked } from '@material-ui/icons';
import useSWR from "swr";
import "./App.css";
import AddTodo from "./components/AddTODO";

export interface Todo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

export const ENDPOINT = "http://localhost:4000";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {
  const { data, mutate } = useSWR<Todo[]>("api/todos", fetcher);

  async function markTodoAsDone(id: number) {
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/done`, {
      method: "PATCH",
    }).then((r) => r.json());

    mutate(updated);
  }

  async function deleteTodo(id: number) {
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/delete`, {
      method: "DELETE",
    }).then((r) => r.json());

    mutate(updated);
  }



  return (
    <Box
      sx={(theme) => ({
        background:"white",
        padding: "2rem",
        width: "100%",
        maxWidth: "40rem",
        margin: "0 auto",
      })}
    >
      
<List spacing="xs" size="sm" mb={12} center>
  {data?.map((todo) => (
    <List.Item key={`todo_list__${todo.id}`}>
      <div>
        {todo.title}
        <Button
          onClick={() => markTodoAsDone(todo.id)}
          startIcon={todo.done ? <CheckCircle /> : <RadioButtonUnchecked />}
        >
          Set as done
        </Button>
        <Button
        onClick={() => deleteTodo(todo.id)}>
          Delete
        </Button>
      </div>
    </List.Item>
  ))}
</List>


      <AddTodo mutate={mutate} />
    </Box>
  );
}

export default App;