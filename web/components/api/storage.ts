import {API_URL} from "@/components/api/config";

export interface Storage {
    id: number;
    name: string;
}

export async function createStorage(storageName: string) {
    await fetch(API_URL + "/api/storage", {
        method: "POST",
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({'storage_name': storageName})
    });
}

export async function deleteStorage(storageID: number) {
    await fetch(API_URL + "/api/storage", {
        method: "DELETE",
        credentials: 'include',
        headers: {
            'storage-id': storageID.toString(),
        }
    });
}

export async function getStorages(): Promise<{ storages: Storage[] }> {
    const resp = await fetch(API_URL + "/api/storage", {
        credentials: 'include',
    });
    if (!resp.ok) {
        return {storages: []};
    }

    return resp.json();
}
