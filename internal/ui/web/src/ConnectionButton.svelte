<script lang="ts">
    import {
        Button,
    } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import type { ButtonColor } from "sveltestrap/src/Button";
    import type { ConnectionItem } from "./Connection";
    import { ConnectionStatus } from "./Connection";

    const dispatch = createEventDispatcher();

    const colorConnection = (item: ConnectionItem): ButtonColor => {
        switch (item.status) {
            case ConnectionStatus.ATTEMPTING:
                return 'secondary';
            case ConnectionStatus.FAILED:
                return 'secondary';
            case ConnectionStatus.CONNECTED:
                return 'danger';
            case ConnectionStatus.DISCONNECTED:
                return 'success';
            default:
                return 'success';
        }
    }

    const btnConnection = (item: ConnectionItem): string => {
        switch (item.status) {
            case ConnectionStatus.ATTEMPTING:
                return 'Cancel';
            case ConnectionStatus.FAILED:
                return 'Cancel';
            case ConnectionStatus.CONNECTED:
                return 'Disconnect';
            case ConnectionStatus.DISCONNECTED:
                return 'Connect';
            default:
                return 'Connect';
        }
    }
    export let connection: ConnectionItem;
    export let classes: string = "";
    $: color = colorConnection(connection);
    $: text = btnConnection(connection);
    const toggleConnection = () => {
        connection.status = ConnectionStatus.ATTEMPTING;
        // connection = {...connection, status: ConnectionStatus.ATTEMPTING, active: true}
        dispatch('toggleConnection', {
            connection: connection
        })
    };
</script>

<Button 
    color={color} 
    class={classes}
    on:click={toggleConnection}>
    {text}
</Button>