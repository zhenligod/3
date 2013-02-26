package kernel

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var madd3_code cu.Function

type madd3_args struct {
	arg_dst  unsafe.Pointer
	arg_src1 unsafe.Pointer
	arg_fac1 float32
	arg_src2 unsafe.Pointer
	arg_fac2 float32
	arg_src3 unsafe.Pointer
	arg_fac3 float32
	arg_N    int
	argptr   [8]unsafe.Pointer
}

// Wrapper for madd3 CUDA kernel, asynchronous.
func K_madd3_async(dst unsafe.Pointer, src1 unsafe.Pointer, fac1 float32, src2 unsafe.Pointer, fac2 float32, src3 unsafe.Pointer, fac3 float32, N int, gridDim, blockDim cu.Dim3, str cu.Stream) {
	if madd3_code == 0 {
		madd3_code = cu.ModuleLoadData(madd3_ptx).GetFunction("madd3")
	}

	var a madd3_args

	a.arg_dst = dst
	a.argptr[0] = unsafe.Pointer(&a.arg_dst)
	a.arg_src1 = src1
	a.argptr[1] = unsafe.Pointer(&a.arg_src1)
	a.arg_fac1 = fac1
	a.argptr[2] = unsafe.Pointer(&a.arg_fac1)
	a.arg_src2 = src2
	a.argptr[3] = unsafe.Pointer(&a.arg_src2)
	a.arg_fac2 = fac2
	a.argptr[4] = unsafe.Pointer(&a.arg_fac2)
	a.arg_src3 = src3
	a.argptr[5] = unsafe.Pointer(&a.arg_src3)
	a.arg_fac3 = fac3
	a.argptr[6] = unsafe.Pointer(&a.arg_fac3)
	a.arg_N = N
	a.argptr[7] = unsafe.Pointer(&a.arg_N)

	args := a.argptr[:]
	cu.LaunchKernel(madd3_code, gridDim.X, gridDim.Y, gridDim.Z, blockDim.X, blockDim.Y, blockDim.Z, 0, str, args)
}

// Wrapper for madd3 CUDA kernel, synchronized.
func K_madd3(dst unsafe.Pointer, src1 unsafe.Pointer, fac1 float32, src2 unsafe.Pointer, fac2 float32, src3 unsafe.Pointer, fac3 float32, N int, gridDim, blockDim cu.Dim3) {
	str := Stream()
	K_madd3_async(dst, src1, fac1, src2, fac2, src3, fac3, N, gridDim, blockDim, str)
	SyncAndRecycle(str)
}

const madd3_ptx = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry madd3(
	.param .u64 madd3_param_0,
	.param .u64 madd3_param_1,
	.param .f32 madd3_param_2,
	.param .u64 madd3_param_3,
	.param .f32 madd3_param_4,
	.param .u64 madd3_param_5,
	.param .f32 madd3_param_6,
	.param .u32 madd3_param_7
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<13>;
	.reg .f32 	%f<21>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd9, [madd3_param_0];
	ld.param.u64 	%rd6, [madd3_param_1];
	ld.param.f32 	%f7, [madd3_param_2];
	ld.param.u64 	%rd7, [madd3_param_3];
	ld.param.f32 	%f8, [madd3_param_4];
	ld.param.u64 	%rd8, [madd3_param_5];
	ld.param.f32 	%f9, [madd3_param_6];
	ld.param.u32 	%r2, [madd3_param_7];
	cvta.to.global.u64 	%rd1, %rd9;
	cvta.to.global.u64 	%rd2, %rd8;
	cvta.to.global.u64 	%rd3, %rd7;
	cvta.to.global.u64 	%rd4, %rd6;
	.loc 2 10 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 12 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_8;

	.loc 2 13 1
	setp.eq.s64 	%p2, %rd6, 0;
	mov.f32 	%f10, 0f3F800000;
	.loc 2 13 1
	mov.f32 	%f20, %f10;
	@%p2 bra 	BB0_3;

	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd4, %rd10;
	ld.global.f32 	%f1, [%rd11];
	mov.f32 	%f20, %f1;

BB0_3:
	.loc 2 13 1
	mov.f32 	%f2, %f20;
	.loc 2 14 1
	setp.eq.s64 	%p3, %rd7, 0;
	mov.f32 	%f19, %f10;
	@%p3 bra 	BB0_5;

	mul.wide.s32 	%rd12, %r1, 4;
	add.s64 	%rd13, %rd3, %rd12;
	ld.global.f32 	%f19, [%rd13];

BB0_5:
	.loc 2 16 1
	cvt.s64.s32 	%rd5, %r1;
	.loc 2 15 1
	setp.eq.s64 	%p4, %rd8, 0;
	mov.f32 	%f18, %f10;
	@%p4 bra 	BB0_7;

	shl.b64 	%rd14, %rd5, 2;
	add.s64 	%rd15, %rd2, %rd14;
	ld.global.f32 	%f18, [%rd15];

BB0_7:
	.loc 2 16 1
	mul.f32 	%f13, %f18, %f9;
	fma.rn.f32 	%f14, %f19, %f8, %f13;
	fma.rn.f32 	%f15, %f2, %f7, %f14;
	mul.wide.s32 	%rd16, %r1, 4;
	add.s64 	%rd17, %rd1, %rd16;
	st.global.f32 	[%rd17], %f15;

BB0_8:
	.loc 2 19 2
	ret;
}


`
