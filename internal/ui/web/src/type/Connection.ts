import type { ConnectionStatus } from "../utils/connection";

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