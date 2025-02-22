<script setup lang="ts">
</script>

<template>
    <header>
        <img src="/soar-logo.svg" alt="Soar Logo" draggable="false" />
    </header>
</template>

<style scoped lang="scss">
@import '../styles/variables.scss';

header {
    margin: 0;
    background-color: $header-bg-color;
    padding: 10px;
    z-index: 100;
    display: flex;

    img {
        width: clamp(80px, 15lvw, 120px);
    }
}
</style>