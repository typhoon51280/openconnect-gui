import type { ButtonColor } from "sveltestrap/src/Button";
import type { Color } from "sveltestrap/src/shared";
import type { ConnectionItem } from "../type/Connection";

export enum ConnectionStatus {
    CONNECTED = "CONNECTED",
    DISCONNECTED = "DISCONNECTED",
    ATTEMPTING = "ATTEMPTING",
    FAILED = "FAILED",
};

export const connectionBtnText = (item: ConnectionItem): string => {
    if (item ===undefined) {
        return '';
    }
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

export const connectionBtnColor = (item: ConnectionItem): ButtonColor => {
    if (item ===undefined) {
        return undefined;
    }
    switch (item.status) {
        case ConnectionStatus.ATTEMPTING:
            return 'secondary';
        case ConnectionStatus.FAILED:
            return 'secondary';
        case ConnectionStatus.CONNECTED:
            return 'danger';
        case ConnectionStatus.DISCONNECTED:
            return 'primary';
        default:
            return 'primary';
    }
}

export const connectionColor = (item: ConnectionItem): Color => {
    if (item ===undefined) {
        return undefined;
    }
    switch (item.status) {
        case ConnectionStatus.ATTEMPTING:
            return 'warning';
        case ConnectionStatus.FAILED:
            return 'danger';
        case ConnectionStatus.CONNECTED:
            return 'success';
        case ConnectionStatus.DISCONNECTED:
            return 'secondary';
        default:
            return 'info';
    }
}