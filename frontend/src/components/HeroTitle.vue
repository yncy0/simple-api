<script setup lang="ts">
interface Props {
  title: string;
}

const props = defineProps<Props>();
</script>

<template>
  <h1>
    {{ props.title }}
  </h1>
</template>

<style scoped>
h1 {
  color: var(--color-primary-500);
  font-size: 3rem;
  font-weight: bold;
  padding: 0;
  margin: 0;
  text-align: center;
}

@media (width >=48rem) {
  h1 {
    font-size: 4rem;
  }
}
</style>
