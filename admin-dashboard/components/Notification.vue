<template>
  <div v-if="visible" class="notification" :class="[type]" @click="hide">
    <p>{{ message }}</p>
  </div>
</template>

<script>
export default {
  props: {
    message: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      default: "success",
    },
    duration: {
      type: Number,
      default: 3000,
    },
  },
  data() {
    return {
      visible: true,
    };
  },
  methods: {
    hide() {
      this.visible = false;
    },
  },
  mounted() {
    if (this.duration > 0) {
      setTimeout(() => {
        this.hide();
      }, this.duration);
    }
  },
};
</script>
