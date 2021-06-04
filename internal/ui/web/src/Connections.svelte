<script lang="ts">
    import { onMount } from "svelte";
    import {
        Button,
        Card,
        CardHeader,
        CardTitle,
        CardBody,
        CardFooter,
        Accordion,
    } from "sveltestrap";
    import type { ConnectionItem } from "./Connection";
    import { ConnectionStatus } from "./Connection";
    import ConnectionModal from './ConnectionModal.svelte';
import ConnectionView from "./ConnectionView.svelte";

    let connections: Array<ConnectionItem> = []

    onMount(async () => {
		connections = await globalThis.Init();
	});
    
    let open;
    let toggle;

    const add = (event) => {
        const connection: ConnectionItem = event.detail.connection
        connection.active = true;
        connections = [...connections, connection];
    }

    const login = (connection: ConnectionItem) => {
		console.log("loginUrl", connection.url);
		globalThis.Login(connection)
		.then((result) => {
			console.log("login result:", result);
            if (result) {
                connection.status = ConnectionStatus.CONNECTED;
            }
		})
	}

</script>

<style></style>

<Card>
    <CardHeader>
        <CardTitle class="d-inline-block">Connections</CardTitle>
        <Button outline color="secondary" class="float-end" on:click={toggle}>Add</Button>
    </CardHeader>
    <CardBody>
        <Accordion>
            {#each connections as connection}
                <ConnectionView {connection}></ConnectionView>
            {/each}
        </Accordion>
    </CardBody>
    <CardFooter>Footer</CardFooter>
    <ConnectionModal bind:open bind:toggle on:add={add}></ConnectionModal>
</Card>