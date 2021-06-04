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