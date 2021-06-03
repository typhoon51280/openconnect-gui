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
        AccordionItem,

    } from "sveltestrap";
    import type { ConnectionItem } from './Connection.svelte';
    import Connection from './Connection.svelte';

    let connections: Array<ConnectionItem> = []

    onMount(async () => {
		connections = await globalThis.Init()
	});
    
    let open = false;
    const toggle = () => (open = !open);

    const add = (event) => {
        const connection: ConnectionItem = event.detail.connection
        connection.active = true
        connections = [...connections, connection]
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
                <AccordionItem active={connection.active} header={connection.name}>
                    <p>{connection.url}</p>
                </AccordionItem>
            {/each}
        </Accordion>
    </CardBody>
    <CardFooter>Footer</CardFooter>
    <Connection {open} {toggle} on:add={add}></Connection>
</Card>