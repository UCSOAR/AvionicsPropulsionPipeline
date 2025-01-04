<script setup lang="ts">
import { ref } from 'vue';
import { devEndpoint } from '../utils/constants';

const selectedFile = ref<File | null>(null);
const uploadError = ref<string | null>(null);

const endpoint = new URL('/upload', devEndpoint);

const onFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    const file: File | null = target.files ? target.files[0] : null;

    selectedFile.value = file;
};

const uploadFile = async () => {
    if (!selectedFile.value) {
        uploadError.value = 'Please select a file to upload.';
        return;
    }

    try {
        const formData = new FormData();
        formData.append('file', selectedFile.value);

        const result = await fetch(endpoint, {
            method: 'POST',
            body: formData,
        });

        if (result.ok) {
            uploadError.value = null;
            selectedFile.value = null;

            alert('File uploaded successfully!');
        } else {
            uploadError.value = result.statusText;
        }
    } catch (err) {
        uploadError.value = err as string;
    }
};
</script>

<template>
    <!-- File Upload Form -->
    <form @submit.prevent="uploadFile" enctype="multipart/form-data">
        <input type="file" @change="onFileChange" />
        <button type="submit">Upload</button>
    </form>
    <!-- Show Current File -->
    <p v-if="selectedFile">Selected File: {{ selectedFile.name }}</p>
    <!-- Error Message -->
    <p v-if="uploadError">{{ uploadError }}</p>
</template>
