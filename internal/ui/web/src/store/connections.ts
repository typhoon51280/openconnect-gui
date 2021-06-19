import { writable } from "svelte/store";
import type { ConnectionItem } from "../type/Connection";

let connections: ConnectionItem[] = [];
let currentConnection: ConnectionItem = {};
let { subscribe, update, set } = writable(connections);
let add = (connection: ConnectionItem) => update(connections => {
    return [...connections, connection];
});
connections.push
let connect = async (connection: ConnectionItem) => {
    const connected = await globalThis.Connect(connection);
    if (connected) {
        connection.active = true;
        currentConnection = connection;
        console.log("Connected: ", currentConnection)
    }
    // set(data);
}
let current = (): ConnectionItem => {
    return currentConnection;
}

export default {
    subscribe,
    update,
    set,
    add,
    connect,
    current,
}