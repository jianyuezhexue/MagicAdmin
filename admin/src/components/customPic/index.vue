<template>
  <span class="headerAvatar">
    <template v-if="picType === 'avatar'">
      <el-avatar v-if="userInfo.headImg" :size="24" :src="avatar" />
      <!-- <el-avatar v-else :size="24" :src="require('@/assets/noBody.png')" /> -->
      <el-avatar v-else :size="24" src="https://www.pubgmobile.com/cp/pubgpic/images/100001.png" />
    </template>
    <template v-if="picType === 'img'">
      <img v-if="userInfo.headImg" :src="avatar" class="avatar">
      <!-- <img v-else :src="require('@/assets/noBody.png')" class="avatar"> -->
      <img v-else src="https://www.pubgmobile.com/cp/pubgpic/images/100001.png" class="avatar">
    </template>
    <template v-if="picType === 'file'">
      <img :src="file" class="file">
    </template>
  </span>
</template>

<script>
import { mapGetters } from "vuex";
const path = import.meta.env.VITE_BASE_API;
export default {
  name: "CustomPic",
  props: {
    picType: {
      type: String,
      required: false,
      default: "avatar",
    },
    picSrc: {
      type: String,
      required: false,
      default: "",
    },
  },
  data() {
    return {
      path: path + "/",
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
    avatar() {
      if (this.picSrc === "") {
        if (
          this.userInfo.headImg !== "" &&
          this.userInfo.headImg.slice(0, 4) === "http"
        ) {
          return this.userInfo.headImg;
        }
        return this.path + this.userInfo.headImg;
      } else {
        if (this.picSrc !== "" && this.picSrc.slice(0, 4) === "http") {
          return this.picSrc;
        }
        return this.path + this.picSrc;
      }
    },
    file() {
      if (this.picSrc && this.picSrc.slice(0, 4) !== "http") {
        return this.path + this.picSrc;
      }
      return this.picSrc;
    },
  },
};
</script>

<style scoped>
.headerAvatar {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-right: 8px;
}
.file {
  width: 80px;
  height: 80px;
  position: relative;
}
</style>
