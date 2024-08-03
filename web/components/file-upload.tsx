'use client';

import {UploadIcon} from "@/components/icons";
import {useCallback, useState} from 'react';
import {decodeString} from "@/components/utils/encoding";

export default function FileUpload({storageName}: { storageName: string }) {
    const [isDragging, setIsDragging] = useState(false);

    const handleDragEnter = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
    }) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(true);
    }, []);

    const handleDragLeave = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
    }) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(false);
    }, []);

    const handleDragOver = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
    }) => {
        e.preventDefault();
        e.stopPropagation();
    }, []);

    const handleDrop = useCallback((e: {
        preventDefault: () => void;
        stopPropagation: () => void;
        dataTransfer: { files: any; };
    }) => {
        e.preventDefault();
        e.stopPropagation();
        setIsDragging(false);

        const files = [...e.dataTransfer.files];
        // Handle the dropped files here
        console.log('Dropped files:', files);
    }, []);

    const handleClick = useCallback(() => {
        const input = document.createElement('input');
        input.type = 'file';
        input.multiple = true;
        input.onchange = (e) => {
            // @ts-ignore
            const files = [...e.target.files];

            for (let i = 0; i < files.length; i++) {
                const file = files[i];

                if (file.size > 5 * 1024 * 1024) {
                    alert('File size exceeds 5MB limit. Only files up to 5MB will be uploaded.');
                    continue;
                }

                console.log(file.name);
            }
        };
        input.click();
    }, []);

    return (
        <>
            <div className="flex items-center justify-between mb-6">
                <h1 className="text-2xl font-bold">Upload files
                    to {decodeString(storageName)}</h1>
            </div>
            <div
                className={`border rounded-lg p-6 flex items-center justify-center h-[100%] bg-muted ${isDragging ? 'border-primary' : ''}`}
                onDragEnter={handleDragEnter}
                onDragOver={handleDragOver}
                onDragLeave={handleDragLeave}
                onDrop={handleDrop}
                onClick={handleClick}
            >
                <div className="text-center space-y-2">
                    <UploadIcon className="w-12 h-12 text-primary mx-auto"/>
                    <p className="text-lg font-medium">Drag and drop files to
                        upload</p>
                    <p className="text-muted-foreground">or click to select
                        files from your device</p>
                </div>
            </div>
        </>
    )
}
