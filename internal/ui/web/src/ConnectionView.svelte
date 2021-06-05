<script lang="ts">
    import { 
        AccordionItem, Badge, Label
    } from "sveltestrap";
    import { connectionColor, ConnectionStatus } from "./utils/connection";
    import type { ConnectionItem } from "./type/Connection";
    import ConnectionButton from "./ConnectionButton.svelte";

    export let connection: ConnectionItem = {};

    $: color = connectionColor(connection);

    const toggleConnection = (event) => {
        let item: ConnectionItem = event.detail.connection;
        console.log('item: ', item);
        console.log('connection: ', connection);
        item.status = ConnectionStatus.ATTEMPTING;
        connection = item
    }

</script> 

<AccordionItem active={connection.active} header={connection.name}>
    <div class="d-flex flex-wrap">
        <div class="text-truncate mt-1"><strong>{connection.url}</strong></div>
        <ConnectionButton classes="ms-auto mt-1" {connection} on:connection:toggle={toggleConnection}></ConnectionButton>
    </div>
    <div>
        <Badge pill {color}>{connection.status || ConnectionStatus.DISCONNECTED}</Badge>
    </div>
</AccordionItem>