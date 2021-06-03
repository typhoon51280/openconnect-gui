<script lang="ts" context="module">
    export class ConnectionItem {
        name: string
        url: string
        username: string
        password: string
        active: boolean
    }
</script>

<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import {
        Button,
        Form,
        FormGroup,
        Label,
        Input,
        Modal,
        ModalHeader,
        ModalBody,
        ModalFooter,
    } from "sveltestrap";
    export let open = false;
    export let toggle = () => (open = !open);

    let connection: ConnectionItem = new ConnectionItem()
    const dispatch = createEventDispatcher()
    
    const save = async (event: Event) =>  {
        event.preventDefault();
        event.stopPropagation();
        const form: any = event.currentTarget;
        if (form.checkValidity()) {
            dispatch('add', {
                connection: connection
            });
            connection = new ConnectionItem();
            toggle();
        }
        form.classList.add('was-validated');
    }
</script>

<Modal isOpen={open} {toggle}>
    <Form class="needs-validation" novalidate on:submit={save}>
        <ModalHeader {toggle}>Modal title</ModalHeader>
        <ModalBody>
            <FormGroup>
                <Label for="name">Name</Label>
                <Input
                    type="text"
                    name="name"
                    id="name"
                    bind:value={connection.name}
                    required
                    placeholder="name" />
            </FormGroup>
            <FormGroup>
                <Label for="url">Url</Label>
                <Input
                    type="url"
                    name="url"
                    id="url"
                    bind:value={connection.url}
                    required
                    placeholder="url" />
            </FormGroup>
            <FormGroup>
                <Label for="username">Username</Label>
                <Input
                    type="text"
                    name="username"
                    id="username"
                    bind:value={connection.username}
                    required
                    placeholder="username" />
            </FormGroup>
            <FormGroup>
                <Label for="password">Password</Label>
                <Input
                    type="password"
                    name="password"
                    id="password"
                    bind:value={connection.password}
                    required
                    placeholder="password" />
            </FormGroup>
        </ModalBody>
        <ModalFooter>
        <Button color="primary" type="submit">Save</Button>
        <!-- <Button color="secondary" on:click={toggle}>Cancel</Button> -->
        </ModalFooter>
    </Form>
</Modal>

