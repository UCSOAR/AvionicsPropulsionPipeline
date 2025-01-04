<script setup lang="ts">
import { ref } from 'vue';

const selectedFile = ref<File | null>(null);
const uploadError = ref<string | null>(null);

const onFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    const file: File | null = target.files ? target.files[0] : null;

    selectedFile.value = file;
};

const uploadFile = () => {
    console.log('Uploading file...');
};
</script>

<template>
    <!-- File Upload Form -->
    <form @submit.prevent="uploadFile">
        <input type="file" @change="onFileChange" />
        <button type="submit">Upload</button>
    </form>
    <!-- Show Current File -->
    <p v-if="selectedFile">Selected File: {{ selectedFile.name }}</p>
    <!-- Error Message -->
    <p v-if="uploadError">{{ uploadError }}</p>
</template>
