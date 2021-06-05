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
        Alert,
    } from "sveltestrap";
    import { connectionColor } from "./utils/connection";
    import type { ConnectionItem } from "./type/Connection";
    import ConnectionModal from './ConnectionModal.svelte';
    import ConnectionView from "./ConnectionView.svelte";
    import connections from "./store/connections";
    
    let open;
    let toggle;
    let status: string;

    $: color = connectionColor($connections.find((item) => item.active));
    $: currentConnection = $connections.find((item) => item.active) || {};

    const add = (event) => {
        const connection: ConnectionItem = event.detail.connection
        connection.active = true;
        connections.add(connection)
    }

    onMount(async () => {
		$connections = await globalThis.Init();
	});

</script>

<style></style>

<Card>
    <CardHeader>
        <CardTitle class="d-inline-block">Connections</CardTitle>
        <Button outline color="info" class="float-end" on:click={toggle}>Add</Button>
    </CardHeader>
    <CardBody>
        <Accordion>
            {#each $connections as connection}
                <ConnectionView {connection}></ConnectionView>
            {/each}
        </Accordion>
    </CardBody>
    <CardFooter>
        <Alert {color}>
            <h5 class="alert-heading text-capitalize">Status:</h5>
            <span>{currentConnection.status || ''}</span>
        </Alert>
    </CardFooter>
    <ConnectionModal bind:open bind:toggle on:add={add}></ConnectionModal>
</Card>