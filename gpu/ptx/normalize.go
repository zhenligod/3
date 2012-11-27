package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["normalize"] = NORMALIZE }

const NORMALIZE = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_0000395c_00000000-9_normalize.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/common_func.h"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"
	.file	4 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/normalize.cu"

.visible .func  (.param .align 4 .b8 func_retval0[12]) _Z9normalize6float3(
	.param .align 4 .b8 _Z9normalize6float3_param_0[12]
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<15>;


	ld.param.f32 	%f3, [_Z9normalize6float3_param_0+8];
	ld.param.f32 	%f1, [_Z9normalize6float3_param_0];
	ld.param.f32 	%f2, [_Z9normalize6float3_param_0+4];
	.loc 2 79 1
	mul.f32 	%f8, %f2, %f2;
	fma.rn.f32 	%f9, %f1, %f1, %f8;
	fma.rn.f32 	%f10, %f3, %f3, %f9;
	.loc 3 991 5
	sqrt.rn.f32 	%f4, %f10;
	mov.f32 	%f14, 0f00000000;
	.loc 2 79 1
	setp.eq.f32 	%p1, %f4, 0f00000000;
	@%p1 bra 	BB0_2;

	rcp.rn.f32 	%f14, %f4;

BB0_2:
	.loc 2 80 1
	mul.f32 	%f11, %f14, %f1;
	mul.f32 	%f12, %f14, %f2;
	mul.f32 	%f13, %f14, %f3;
	st.param.f32	[func_retval0+0], %f11;
	st.param.f32	[func_retval0+4], %f12;
	st.param.f32	[func_retval0+8], %f13;
	ret;
}

.visible .entry normalize(
	.param .u64 normalize_param_0,
	.param .u64 normalize_param_1,
	.param .u64 normalize_param_2,
	.param .u32 normalize_param_3
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<4>;
	.reg .s64 	%rd<11>;


	ld.param.u64 	%rd4, [normalize_param_0];
	ld.param.u64 	%rd5, [normalize_param_1];
	ld.param.u64 	%rd6, [normalize_param_2];
	ld.param.u32 	%r2, [normalize_param_3];
	cvta.to.global.u64 	%rd1, %rd6;
	cvta.to.global.u64 	%rd2, %rd5;
	cvta.to.global.u64 	%rd3, %rd4;
	.loc 4 6 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 4 7 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB1_2;

	.loc 4 9 1
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd3, %rd7;
	add.s64 	%rd9, %rd2, %rd7;
	ld.global.f32 	%f1, [%rd9];
	add.s64 	%rd10, %rd1, %rd7;
	ld.global.f32 	%f2, [%rd10];
	ld.global.f32 	%f3, [%rd8];
	.loc 4 11 1
	st.global.f32 	[%rd8], %f3;
	.loc 4 12 1
	st.global.f32 	[%rd9], %f1;
	.loc 4 13 1
	st.global.f32 	[%rd10], %f2;

BB1_2:
	.loc 4 15 2
	ret;
}


`
