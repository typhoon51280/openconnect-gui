<script lang="ts">
    import {
        Button,
    } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import { connectionBtnColor, connectionBtnText } from "./utils/connection";
    import type { ConnectionItem } from "./type/Connection";

    export let connection: ConnectionItem;
    export let classes: string = "";

    const dispatch = createEventDispatcher();

    $: color = connectionBtnColor(connection);
    $: text = connectionBtnText(connection);

    const toggleConnection = () => {
        dispatch('connection:toggle', {
            connection: connection,
        })
    };

</script>

<div class={classes}>
<Button 
    color={color}
    on:click={toggleConnection}>
    {text}
</Button>
</div>