import { writable } from "svelte/store";
import type { ConnectionItem } from "../type/Connection";

let connections: ConnectionItem[] = [];
let { subscribe, update, set } = writable(connections);
let add = (connection: ConnectionItem) => update(connections => {
    return [...connections, connection];
});
let connect = async (connection: ConnectionItem) => {
    const data: ConnectionItem[] = await globalThis.Connect(connection);
    set(data);
}

export default {
    subscribe,
    update,
    set,
    add,
    connect,
}