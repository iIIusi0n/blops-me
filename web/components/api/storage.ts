import {cookies} from "next/headers";
import {INTERNAL_API_URL} from "@/components/api/config";

export interface Storage {
    id: number;
    name: string;
}

export async function createStorage(storageName: string) {
    const token = cookies().get('token')?.value;
    await fetch( `${INTERNAL_API_URL}/api/storage`, {
        method: "POST",
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            'Cookie': `token=${token}`
        },
        body: JSON.stringify({'storage_name': storageName})
    });
}

export async function deleteStorage(storageID: number) {
    const token = cookies().get('token')?.value;
    await fetch( `${INTERNAL_API_URL}/api/storage`, {
        method: "DELETE",
        credentials: 'include',
        headers: {
            'storage-id': storageID.toString(),
            'Cookie': `token=${token}`
        }
    });
}

export async function getStorages(): Promise<{ storages: Storage[] }> {
    const token = cookies().get('token')?.value;
    const resp = await fetch( `${INTERNAL_API_URL}/api/storage`, {
        credentials: 'include',
        headers: {
            'Cookie': `token=${token}`
        }
    });
    if (!resp.ok) {
        return {storages: []};
    }

    return resp.json();
}
