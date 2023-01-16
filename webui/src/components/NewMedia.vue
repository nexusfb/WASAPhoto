<script>
export default {
	name: 'NewMedia',
    props: {
        pic: { 
            type: String,
            default:"1a9bc88e-55a3-41ed-99d0-873ab630507e.jpg"
        },
        caption:{
            type: String,
            default: ""
        }
    }
}
</script>

<template>
	<div class="alert alert-danger" role="alert" >
        <div class="container">
          <form @submit.prevent="onSubmit">
              <div class="form-group">
                  <input type="file" @change="handleImageUpload">
              </div>
              <img id="preview-image" v-if="preview" :src="preview" :width="400" :height="400">
              <div class="form-group">
                  <input type="text" v-model="caption" placeholder="Write a caption here" class="form-control">
              </div>
              <div class="form-group">
                  <button class="btn btn-primary btn-block btn-lg">Upload File</button>
              </div>
          </form>
      </div>    
		<small v-if="details"><br /><br />Details: {{ details }}</small>
	</div>
</template>

<style>
.alert.alert-danger {
    height: 400px;
    padding-left: 10px;
    padding-right: 16px;
	padding-top: 10px;
    background-color:rgb(255, 255, 255);
    border-color: beige;
    border-radius: 20px;

}
</style>
