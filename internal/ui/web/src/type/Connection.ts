import type { ButtonColor } from "sveltestrap/src/Button";
import type { Color } from "sveltestrap/src/shared";

export enum ConnectionStatus {
    CONNECTED = "CONNECTED",
    DISCONNECTED = "DISCONNECTED",
    ATTEMPTING = "ATTEMPTING",
    FAILED = "FAILED",
};

export type LoginProfile = {
    name?: string
    username?: string
    password?: string
};

export type ConnectionItem = {
    name?: string
    url?: string
    active?: boolean
    autologin?: boolean
    login?: LoginProfile
    status?: ConnectionStatus
};

export const connectionBtnText = (item: ConnectionItem): string => {
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

export const connectionColor = (item: ConnectionItem): Color => {
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
            return 'secondary';
    }
}