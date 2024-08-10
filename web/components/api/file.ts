import {cookies} from "next/headers";
import {getStorages} from "@/components/api/storage";

export interface File {
    id: number;
    name: string;
    type: string;
    last_modified: string;
    size: number;
}

export async function resolveStorageID(storageName: string): Promise<number> {
    const data = await getStorages();
    const storage = data.storages.find(storage => storage.name === storageName);
    return storage?.id ?? -1;
}

export async function getFiles(storageId: number, pathId: number): Promise<File[]> {
    const token = cookies().get('token')?.value;
    const resp = await fetch(`${process.env.APP_API_URL}/api/storage/${storageId}/file${pathId ? `?path=${pathId}` : ''}`, {
        credentials: 'include',
        headers: {
            'Cookie': `token=${token}`
        }
    });
    if (!resp.ok) {
        return [];
    }

    return resp.json();
}

export async function getPath(storageId: number, pathId: number): Promise<any> {
    const token = cookies().get('token')?.value;
    const resp = await fetch(`${process.env.APP_API_URL}/api/storage/${storageId}/path/${pathId}`, {
        credentials: 'include',
        headers: {
            'Cookie': `token=${token}`
        }
    });
    if (!resp.ok) {
        return '';
    }

    return resp.json();
}

export async function getParentID(storageId: number, pathId: number): Promise<number> {
    const token = cookies().get('token')?.value;
    const resp = await fetch(`${process.env.APP_API_URL}/api/storage/${storageId}/parent/${pathId}`, {
        credentials: 'include',
        headers: {
            'Cookie': `token=${token}`
        }
    });
    if (!resp.ok) {
        return -1;
    }

    return resp.json();
}
