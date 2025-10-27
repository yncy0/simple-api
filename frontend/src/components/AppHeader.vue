<script setup lang="ts">
import { Icon } from "@iconify/vue";
</script>

<template>
  <header>
    <Icon icon="lucide:braces" class="icon" />

    <a href="https://github.com/yncy0/simple-api" target="_blank">
      <Icon icon="simple-icons:github" class="github-icon" />
    </a>
  </header>
</template>

<style scoped>
header {
  font-size: 1.75rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: fixed;
  top: 0px;
  left: 0px;
  width: 100vw;
  padding-inline: 2rem;
  padding-top: 1rem;
  padding-bottom: 1rem;
}

@media (width >=48rem) {
  header {
    font-size: 2rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: fixed;
    padding-inline: 5rem;
  }
}

.icon {
  color: var(--color-primary-500);
}

.github-icon {
  color: var(--color-primary-500);
}
</style>
