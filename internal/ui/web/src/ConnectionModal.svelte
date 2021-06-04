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
        InputGroup,
        InputGroupText,
    } from "sveltestrap";
    import type { ConnectionItem, LoginProfile } from "./Connection";
    
    export let open = false;
    export let toggle = () => (open = !open);

    let connection: ConnectionItem = {};
    const dispatch = createEventDispatcher();

    const emptyProfile: LoginProfile = {};

    let profiles: Array<LoginProfile> = [
        emptyProfile,
        {
            name: "DEMO1",
            username: "NTT3DR7",
            password: "blabla"
        },
        {
            name: "DEMO2",
            username: "NTT3DR8",
            password: "blabla"
        },
    ]

    const autologinChanged = () => {
        if (connection.autologin) {
            connection.login = emptyProfile;
        }
    }
    
    const save = async (event: Event) =>  {
        event.preventDefault();
        event.stopPropagation();
        const form: any = event.currentTarget;
        if (form.checkValidity()) {
            dispatch('add', {
                connection: connection
            });
            connection = undefined;
            toggle();
        }
        form.classList.add('was-validated');
    }
</script>

<Modal isOpen={open} {toggle}>
    <Form class="needs-validation" novalidate on:submit={save}>
        <ModalHeader {toggle}>Connection</ModalHeader>
        <ModalBody>
            <FormGroup>
                <Label for="name">Name</Label>
                <Input
                    type="text"
                    name="name"
                    bind:value={connection.name}
                    required
                    placeholder="name" />
            </FormGroup>
            <FormGroup>
                <Label for="url">Url</Label>
                <Input
                    type="url"
                    name="url"
                    bind:value={connection.url}
                    required
                    placeholder="url" />
            </FormGroup>
            <FormGroup>
                <InputGroup>
                    <InputGroupText>
                        <Input
                            type="checkbox"
                            name="autologin"
                            addon
                            bind:checked={connection.autologin}
                            on:change={autologinChanged}
                        />
                    </InputGroupText>
                    <Input
                        name="login"
                        type="select"
                        required
                        bind:value={connection.login}
                        disabled={!connection.autologin}
                    >
                        {#each profiles as profile}
                        <option value={profile}>
                            {profile.name || ''}
                        </option>
                        {/each}
                    </Input>
                  </InputGroup>
            </FormGroup>
            <!-- <FormGroup>
                <Label for="username">Username</Label>
                <Input
                    type="text"
                    name="username"
                    value={connection.login.username}
                    required
                    placeholder="username" />
            </FormGroup>
            <FormGroup>
                <Label for="password">Password</Label>
                <Password
                    bind:value={connection.login.password}
                >
                </Password>
            </FormGroup> -->
        </ModalBody>
        <ModalFooter>
        <Button color="primary" type="submit">Save</Button>
        <Button color="secondary" on:click={toggle}>Cancel</Button>
        </ModalFooter>
    </Form>
</Modal>

