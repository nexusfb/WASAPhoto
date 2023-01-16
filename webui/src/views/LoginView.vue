<script>
export default {
    data: function() {
        return {
            errormsg: true,
            detailedmsg: null,
            loading: false,
            User: {
                UserID: null,
                Username: null,

            }
        }
    },
    methods: {
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/session/", {
                    username: this.Username,
                });
                this.UserID  = response.data,
                localStorage.setItem('Authorization', this.UserID),
                this.$router.push({ path: '/users/'+this.Username })

            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }


    }
}
</script>

<template>
    <div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div
            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">Login</h1>
        </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <div class="mb-3">
            <label for="description" class="form-label">Username</label>
            <input type="string" class="form-control" id="username" v-model="Username" placeholder="please insert username">
        </div>

        <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="LoginUser">
                Login
            </button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </div>
    </div>

	<transition name="popup-t">
    <div
      v-show="active"
      ref="con"
      :class="[`vs-popup-${color}`,{'fullscreen':fullscreen}]"
      class="vs-component con-vs-popup"
      @click="close($event,true)">
      <div
        :style="styleCon"
        class="vs-popup--background"/>
      <div
        ref="popupx"
        :style="stylePopup"
        class="vs-popup">

        <!-- //header -->
        <header
          :style="styleHeader"
          class="vs-popup--header">
          <div class="vs-popup--title">
            <h3>{{ title }}</h3>
            <slot name="subtitle" />
          </div>
          <vs-icon
            v-if="!buttonCloseHidden"
            ref="btnclose"
            :icon-pack="iconPack"
            :icon="iconClose"
            :style="stylePopup"
            class="vs-popup--close vs-popup--close--icon"
            @click="close"
          />
        </header>

        <!-- // slots  -->
        <div
          :style="styleContent"
          :class="classContent"
          class="vs-popup--content">
          <slot/>
        </div>
      </div>
    </div>
  </transition>

</template>

<style>
</style>