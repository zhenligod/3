package kernel

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var kernmulRSymm2Dx_code cu.Function

type kernmulRSymm2Dx_args struct {
	arg_fftMx  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_N1     int
	arg_N2     int
	argptr     [4]unsafe.Pointer
}

// Wrapper for kernmulRSymm2Dx CUDA kernel, asynchronous.
func K_kernmulRSymm2Dx_async(fftMx unsafe.Pointer, fftKxx unsafe.Pointer, N1 int, N2 int, gridDim, blockDim cu.Dim3, str cu.Stream) {
	if kernmulRSymm2Dx_code == 0 {
		kernmulRSymm2Dx_code = cu.ModuleLoadData(kernmulRSymm2Dx_ptx).GetFunction("kernmulRSymm2Dx")
	}

	var a kernmulRSymm2Dx_args

	a.arg_fftMx = fftMx
	a.argptr[0] = unsafe.Pointer(&a.arg_fftMx)
	a.arg_fftKxx = fftKxx
	a.argptr[1] = unsafe.Pointer(&a.arg_fftKxx)
	a.arg_N1 = N1
	a.argptr[2] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[3] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(kernmulRSymm2Dx_code, gridDim.X, gridDim.Y, gridDim.Z, blockDim.X, blockDim.Y, blockDim.Z, 0, str, args)
}

// Wrapper for kernmulRSymm2Dx CUDA kernel, synchronized.
func K_kernmulRSymm2Dx(fftMx unsafe.Pointer, fftKxx unsafe.Pointer, N1 int, N2 int, gridDim, blockDim cu.Dim3) {
	str := Stream()
	K_kernmulRSymm2Dx_async(fftMx, fftKxx, N1, N2, gridDim, blockDim, str)
	SyncAndRecycle(str)
}

const kernmulRSymm2Dx_ptx = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry kernmulRSymm2Dx(
	.param .u64 kernmulRSymm2Dx_param_0,
	.param .u64 kernmulRSymm2Dx_param_1,
	.param .u32 kernmulRSymm2Dx_param_2,
	.param .u32 kernmulRSymm2Dx_param_3
)
{
	.reg .pred 	%p<5>;
	.reg .s32 	%r<26>;
	.reg .f32 	%f<6>;
	.reg .s64 	%rd<11>;


	ld.param.u64 	%rd3, [kernmulRSymm2Dx_param_0];
	ld.param.u64 	%rd4, [kernmulRSymm2Dx_param_1];
	ld.param.u32 	%r3, [kernmulRSymm2Dx_param_2];
	ld.param.u32 	%r4, [kernmulRSymm2Dx_param_3];
	cvta.to.global.u64 	%rd1, %rd4;
	cvta.to.global.u64 	%rd2, %rd3;
	.loc 2 19 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	.loc 2 20 1
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %ctaid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	.loc 2 22 1
	setp.ge.s32 	%p1, %r2, %r4;
	setp.ge.s32 	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	.loc 2 26 1
	mad.lo.s32 	%r11, %r1, %r4, %r2;
	.loc 2 27 1
	sub.s32 	%r12, %r3, %r1;
	mad.lo.s32 	%r13, %r12, %r4, %r2;
	.loc 2 29 1
	shl.b32 	%r14, %r11, 1;
	.loc 2 31 1
	mul.wide.s32 	%rd5, %r14, 4;
	add.s64 	%rd6, %rd2, %rd5;
	.loc 2 32 1
	add.s32 	%r15, %r14, 1;
	mul.wide.s32 	%rd7, %r15, 4;
	add.s64 	%rd8, %rd2, %rd7;
	ld.global.f32 	%f1, [%rd8];
	.loc 2 35 1
	shr.u32 	%r17, %r3, 31;
	add.s32 	%r18, %r3, %r17;
	shr.s32 	%r19, %r18, 1;
	add.s32 	%r20, %r19, 1;
	setp.lt.s32 	%p4, %r1, %r20;
	.loc 2 36 1
	selp.b32 	%r21, %r11, %r13, %p4;
	mul.wide.s32 	%rd9, %r21, 4;
	add.s64 	%rd10, %rd1, %rd9;
	.loc 2 41 1
	ld.global.f32 	%f2, [%rd10];
	.loc 2 31 1
	ld.global.f32 	%f3, [%rd6];
	.loc 2 41 1
	mul.f32 	%f4, %f3, %f2;
	st.global.f32 	[%rd6], %f4;
	.loc 2 42 1
	mul.f32 	%f5, %f1, %f2;
	st.global.f32 	[%rd8], %f5;

BB0_2:
	.loc 2 43 2
	ret;
}


`
