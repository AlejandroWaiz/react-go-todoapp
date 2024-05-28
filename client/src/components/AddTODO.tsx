import {useState} from "react"
import {useForm} from "@mantine/hooks"
import {Button, Modal, Group, TextInput, Textarea} from "@mantine/core"
import { ENDPOINT, Todo } from "../App"
import { KeyedMutator } from "swr"

function AddTodo({mutate}: {mutate: KeyedMutator<Todo[]>}) {

    const [open, setOpen ] = useState(Boolean)

    const form = useForm({
        initialValues: {
            title:"",
            body:"",
        }
    })

    async function createTODO(values: {title:string, body:string}){

            const updated = await fetch(`${ENDPOINT}/api/todos`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body:JSON.stringify(values),
            }).then((response) => response.json())

            mutate(updated)
            form.reset()
            setOpen(false)

    }

    return <>
    
        <Modal opened={open} onClose={()=> setOpen(false)} title="Add todo">

            <form onSubmit={form.onSubmit(createTODO)}>
                <TextInput required mb={12} label="Todo" placeholder="What do you want to do?" {...form.getInputProps("title")}/>
                <Textarea required mb={12} label="Body" placeholder="Tell me more..." {...form.getInputProps("body")}></Textarea>
                <Button type="submit">Create TODO</Button>
            </form>


        </Modal>

        <Group position="center">

            <Button fullWidth mb={12} onClick={() => setOpen(true)}>
                ADD TODO
            </Button>

        </Group>
    
        </>

}

export default AddTodo;