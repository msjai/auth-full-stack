<template>
  <div class="px-4">
    <h1 class="text-xl text-center font-semibold mb-2">
      {{ isLogin ? 'Login' : 'Sign Up' }}
    </h1>
    <div class="relative group">
    <input
        type="text"
        id="Email Address"
        placeholder="Email Address"
        @input="emailField.handleChange"
        @blur="emailField.handleBlur"
        class="px-4 my-2 min-w-full mx-auto border border-gray-500 rounded-full focus:outline-none focus:ring-1 focus:border-blue-300"
    />
    </div>
    <p
        v-show="emailField.meta.touched && !emailField.meta.valid"
        class="text-center text-red-500"
    >
      {{ emailField.errorMessage || 'Field is Required' }}
    </p>
    <input
        type="password"
        placeholder="Password"
        @input="
        passwordField.handleChange($event), confirmPasswordField.validate()
      "
        @blur="passwordField.handleBlur"
        class="px-4 my-2 min-w-full mx-auto border border-gray-500 rounded-full focus:outline-none focus:ring-1 focus:border-blue-300"
    />
    <p
        v-show="passwordField.meta.touched && !passwordField.meta.valid"
        class="text-center text-red-500"
    >
      {{ passwordField.errorMessage || 'Field is Required' }}
    </p>

    <template v-if="!isLogin">
      <input
          type="password"
          placeholder="Confirm Password"
          @input="confirmPasswordField.handleChange"
          @blur="confirmPasswordField.handleBlur"
          class="px-4 my-2 min-w-full mx-auto border border-gray-500 rounded-full focus:outline-none focus:ring-1 focus:border-blue-300"
      />
      <p
          v-show="confirmPasswordField.meta.touched && !confirmPasswordField.meta.valid"
          class="text-center text-red-500"
      >
        {{ confirmPasswordField.errorMessage || 'Field is Required' }}
      </p>
    </template>

    <div class="flex justify-center mt-4">
      <button
          class="btn btn-blue mx-1 flex items-center"
          :disabled="!formMeta.valid"
          @click="submitForm"
      >
        <span>{{ isLogin ? 'Login' : 'Sign Up' }}</span>
        <Loader
            v-if="isSubmitting"
            class="animate-spin stroke-current text-white ml-2"
            :height="16"
        />
      </button>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive, computed, watch } from 'vue';
import { useField, useForm } from 'vee-validate';
import Loader from './ui/Loader.vue';
export default defineComponent({
  name: 'LoginForm',
  components: {
    Loader,
  },
  props: {
    isLogin: {
      type: Boolean,
      default: true,
    },
    isSubmitting: {
      type: Boolean,
      default: false,
    },
  },
  emits: {
    submitAuth: null, // null means we will not validate event
  },
  setup(props, { emit }) {
    const { meta: formMeta, handleSubmit } = useForm();
    const emailField = reactive(useField('email', 'email'));
    const passwordField = reactive(useField('password', 'password'));
    const confirmPasswordValidator = computed(() => {
      return !props.isLogin ? 'confirmPassword:password' : () => true;
    });
    const confirmPasswordField = reactive(
        useField('confirmPassword', confirmPasswordValidator)
    );
    watch(
        () => props.isLogin,
        () => {
          confirmPasswordField.validate();
        }
    );
    const submitForm = handleSubmit((formValues) => {
      emit('submitAuth', {
        email: formValues.email,
        password: formValues.password,
      });
    });
    return {
      emailField,
      passwordField,
      confirmPasswordField,
      submitForm,
      formMeta,
    };
  },
});
</script>