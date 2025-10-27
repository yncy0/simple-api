<script setup lang="ts">
import { Icon } from "@iconify/vue";
</script>

<template>
  <header>
    <Icon icon="lucide:braces" />

    <a href="https://github.com/yncy0/simple-api" target="_blank">
      <Icon icon="simple-icons:github" />
    </a>
  </header>
</template>

<style scoped>
header {
  font-size: 2rem;
  display: flex;
  justify-content: space-between;
}
</style>
