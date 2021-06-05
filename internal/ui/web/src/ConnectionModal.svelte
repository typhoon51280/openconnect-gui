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
    import { ConnectionStatus } from "./type/Connection";
    import type { ConnectionItem, LoginProfile } from "./type/Connection";

    export let open = false;
    export let toggle = () => {
        formClass = formClass.filter((item) => item !== 'was-validated');
        connection = {status: ConnectionStatus.DISCONNECTED};
        open = !open;
    }
    
    let formClass = ['needs-validation'];
    let connection: ConnectionItem = {};
    const emptyProfile: LoginProfile = {};
    const dispatch = createEventDispatcher();

    $: connection.login = connection.autologin ? connection.login : emptyProfile;

    const save = async (event: Event) =>  {
        event.preventDefault();
        event.stopPropagation();
        formClass = [...formClass, 'was-validated']
        const form: HTMLFormElement = event.target as HTMLFormElement;
        if (form.checkValidity()) {
            dispatch('add', {
                connection: connection
            });
            toggle();
        }
    }

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
    
</script>

<Modal isOpen={open} {toggle}>
    <Form class={formClass} novalidate on:submit={save}>
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
        </ModalBody>
        <ModalFooter>
        <Button color="primary" type="submit">Save</Button>
        <Button color="secondary" on:click={toggle}>Cancel</Button>
        </ModalFooter>
    </Form>
</Modal>

