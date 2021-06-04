<script lang="ts">
import { 
    AccordionItem
} from "sveltestrap";
import { ConnectionStatus } from "./Connection";
import type { ConnectionItem } from "./Connection";
import ConnectionButton from "./ConnectionButton.svelte";

const toggleConnection = (event) => {
    let item: ConnectionItem = event.detail.connection;
    console.log('item: ', item);
    console.log('connection: ', connection);
    connection = item
    // connections[connections.findIndex(item => item.name === connection.name)] = connection;
}

export let connection: ConnectionItem;
</script> 

<AccordionItem active={connection.active} header={connection.name}>
    <div class="d-flex flex-wrap">
        <div class="text-truncate">{connection.url}</div>
        <ConnectionButton classes="ms-auto" {connection} on:toggleConnection={toggleConnection}></ConnectionButton>
    </div>
    <div>
        {connection.status || ConnectionStatus.DISCONNECTED}
    </div>
</AccordionItem>