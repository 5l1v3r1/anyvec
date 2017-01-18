//+build !nocuda

package cuda

var kernelPTX = `
//
// Generated by NVIDIA NVVM Compiler
//
// Compiler Build ID: CL-21313162
// Cuda compilation tools, release 8.0, V8.0.53
// Based on LLVM 3.4svn
//

.version 5.0
.target sm_20
.address_size 64

	// .globl	divElements
.const .align 4 .b8 __cudart_i2opi_f[24] = {65, 144, 67, 60, 153, 149, 98, 219, 192, 221, 52, 245, 209, 87, 39, 252, 41, 21, 68, 78, 110, 131, 249, 162};

.visible .entry divElements(
	.param .u64 divElements_param_0,
	.param .u64 divElements_param_1,
	.param .u64 divElements_param_2
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<4>;
	.reg .b32 	%r<5>;
	.reg .b64 	%rd<10>;


	ld.param.u64 	%rd2, [divElements_param_0];
	ld.param.u64 	%rd3, [divElements_param_1];
	ld.param.u64 	%rd4, [divElements_param_2];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd4;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd5, %rd2;
	shl.b64 	%rd6, %rd1, 2;
	add.s64 	%rd7, %rd5, %rd6;
	cvta.to.global.u64 	%rd8, %rd3;
	add.s64 	%rd9, %rd8, %rd6;
	ld.global.f32 	%f1, [%rd9];
	ld.global.f32 	%f2, [%rd7];
	div.rn.f32 	%f3, %f2, %f1;
	st.global.f32 	[%rd7], %f3;

BB0_2:
	ret;
}

	// .globl	expElements
.visible .entry expElements(
	.param .u64 expElements_param_0,
	.param .u64 expElements_param_1
)
{
	.reg .pred 	%p<4>;
	.reg .f32 	%f<15>;
	.reg .b32 	%r<5>;
	.reg .b64 	%rd<7>;


	ld.param.u64 	%rd2, [expElements_param_0];
	ld.param.u64 	%rd3, [expElements_param_1];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd3;
	@%p1 bra 	BB1_2;

	cvta.to.global.u64 	%rd4, %rd2;
	shl.b64 	%rd5, %rd1, 2;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, 0f3FB8AA3B;
	cvt.rzi.f32.f32	%f5, %f4;
	mov.f32 	%f6, 0fBF317200;
	fma.rn.f32 	%f7, %f5, %f6, %f3;
	mov.f32 	%f8, 0fB5BFBE8E;
	fma.rn.f32 	%f9, %f5, %f8, %f7;
	mul.f32 	%f2, %f9, 0f3FB8AA3B;
	// inline asm
	ex2.approx.ftz.f32 %f1,%f2;
	// inline asm
	add.f32 	%f10, %f5, 0f00000000;
	ex2.approx.f32 	%f11, %f10;
	mul.f32 	%f12, %f1, %f11;
	setp.lt.f32	%p2, %f3, 0fC2D20000;
	selp.f32	%f13, 0f00000000, %f12, %p2;
	setp.gt.f32	%p3, %f3, 0f42D20000;
	selp.f32	%f14, 0f7F800000, %f13, %p3;
	st.global.f32 	[%rd6], %f14;

BB1_2:
	ret;
}

	// .globl	tanhElements
.visible .entry tanhElements(
	.param .u64 tanhElements_param_0,
	.param .u64 tanhElements_param_1
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<33>;
	.reg .b32 	%r<10>;
	.reg .b64 	%rd<7>;


	ld.param.u64 	%rd3, [tanhElements_param_0];
	ld.param.u64 	%rd4, [tanhElements_param_1];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd4;
	@%p1 bra 	BB2_5;

	cvta.to.global.u64 	%rd5, %rd3;
	shl.b64 	%rd6, %rd1, 2;
	add.s64 	%rd2, %rd5, %rd6;
	ld.global.f32 	%f1, [%rd2];
	abs.f32 	%f2, %f1;
	setp.ltu.f32	%p2, %f2, 0f3F0CCCCD;
	@%p2 bra 	BB2_3;
	bra.uni 	BB2_2;

BB2_3:
	mul.f32 	%f21, %f1, %f1;
	mov.f32 	%f22, 0fBD57BE66;
	mov.f32 	%f23, 0f3C86A81B;
	fma.rn.f32 	%f24, %f23, %f21, %f22;
	mov.f32 	%f25, 0f3E08677B;
	fma.rn.f32 	%f26, %f24, %f21, %f25;
	mov.f32 	%f27, 0fBEAAAA29;
	fma.rn.f32 	%f28, %f26, %f21, %f27;
	mul.f32 	%f29, %f21, %f28;
	fma.rn.f32 	%f30, %f29, %f1, %f1;
	add.f32 	%f31, %f1, %f1;
	setp.eq.f32	%p4, %f1, 0f00000000;
	selp.f32	%f32, %f31, %f30, %p4;
	bra.uni 	BB2_4;

BB2_2:
	add.f32 	%f10, %f2, %f2;
	mul.f32 	%f11, %f10, 0f3FB8AA3B;
	cvt.rzi.f32.f32	%f12, %f11;
	mov.f32 	%f13, 0fBF317200;
	fma.rn.f32 	%f14, %f12, %f13, %f10;
	mov.f32 	%f15, 0fB5BFBE8E;
	fma.rn.f32 	%f16, %f12, %f15, %f14;
	mul.f32 	%f7, %f16, 0f3FB8AA3B;
	// inline asm
	ex2.approx.ftz.f32 %f6,%f7;
	// inline asm
	ex2.approx.f32 	%f17, %f12;
	mov.f32 	%f18, 0f3F800000;
	fma.rn.f32 	%f9, %f6, %f17, %f18;
	// inline asm
	rcp.approx.ftz.f32 %f8,%f9;
	// inline asm
	mov.f32 	%f19, 0fC0000000;
	fma.rn.f32 	%f20, %f8, %f19, %f18;
	mov.b32 	 %r5, %f20;
	setp.ltu.f32	%p3, %f2, 0f42B00000;
	selp.b32	%r6, %r5, 1065353216, %p3;
	mov.b32 	 %r7, %f1;
	and.b32  	%r8, %r7, -2147483648;
	or.b32  	%r9, %r6, %r8;
	mov.b32 	 %f32, %r9;

BB2_4:
	st.global.f32 	[%rd2], %f32;

BB2_5:
	ret;
}

	// .globl	sinElements
.visible .entry sinElements(
	.param .u64 sinElements_param_0,
	.param .u64 sinElements_param_1
)
{
	.local .align 4 .b8 	__local_depot3[28];
	.reg .b64 	%SP;
	.reg .b64 	%SPL;
	.reg .pred 	%p<15>;
	.reg .f32 	%f<48>;
	.reg .b32 	%r<96>;
	.reg .b64 	%rd<19>;


	mov.u64 	%rd18, __local_depot3;
	cvta.local.u64 	%SP, %rd18;
	ld.param.u64 	%rd9, [sinElements_param_0];
	ld.param.u64 	%rd10, [sinElements_param_1];
	mov.u32 	%r36, %ntid.x;
	mov.u32 	%r37, %ctaid.x;
	mov.u32 	%r38, %tid.x;
	mad.lo.s32 	%r39, %r36, %r37, %r38;
	cvt.u64.u32	%rd1, %r39;
	setp.ge.u64	%p1, %rd1, %rd10;
	@%p1 bra 	BB3_24;

	cvta.to.global.u64 	%rd11, %rd9;
	shl.b64 	%rd12, %rd1, 2;
	add.s64 	%rd2, %rd11, %rd12;
	add.u64 	%rd13, %SP, 0;
	cvta.to.local.u64 	%rd3, %rd13;
	ld.global.f32 	%f43, [%rd2];
	abs.f32 	%f19, %f43;
	setp.neu.f32	%p2, %f19, 0f7F800000;
	@%p2 bra 	BB3_3;

	mov.f32 	%f20, 0f00000000;
	mul.rn.f32 	%f43, %f43, %f20;

BB3_3:
	mul.f32 	%f21, %f43, 0f3F22F983;
	cvt.rni.s32.f32	%r95, %f21;
	cvt.rn.f32.s32	%f22, %r95;
	neg.f32 	%f23, %f22;
	mov.f32 	%f24, 0f3FC90FDA;
	fma.rn.f32 	%f25, %f23, %f24, %f43;
	mov.f32 	%f26, 0f33A22168;
	fma.rn.f32 	%f27, %f23, %f26, %f25;
	mov.f32 	%f28, 0f27C234C5;
	fma.rn.f32 	%f44, %f23, %f28, %f27;
	abs.f32 	%f29, %f43;
	setp.leu.f32	%p3, %f29, 0f47CE4780;
	@%p3 bra 	BB3_13;

	mov.b32 	 %r2, %f43;
	shr.u32 	%r3, %r2, 23;
	shl.b32 	%r42, %r2, 8;
	or.b32  	%r4, %r42, -2147483648;
	mov.u32 	%r87, 0;
	mov.u64 	%rd16, __cudart_i2opi_f;
	mov.u32 	%r86, -6;
	mov.u64 	%rd17, %rd3;

BB3_5:
	.pragma "nounroll";
	mov.u64 	%rd5, %rd17;
	ld.const.u32 	%r45, [%rd16];
	// inline asm
	{
	mad.lo.cc.u32   %r43, %r45, %r4, %r87;
	madc.hi.u32     %r87, %r45, %r4,  0;
	}
	// inline asm
	st.local.u32 	[%rd5], %r43;
	add.s64 	%rd6, %rd5, 4;
	add.s64 	%rd16, %rd16, 4;
	add.s32 	%r86, %r86, 1;
	setp.ne.s32	%p4, %r86, 0;
	mov.u64 	%rd17, %rd6;
	@%p4 bra 	BB3_5;

	and.b32  	%r48, %r3, 255;
	add.s32 	%r49, %r48, -128;
	shr.u32 	%r50, %r49, 5;
	and.b32  	%r9, %r2, -2147483648;
	st.local.u32 	[%rd3+24], %r87;
	mov.u32 	%r51, 6;
	sub.s32 	%r52, %r51, %r50;
	mul.wide.s32 	%rd15, %r52, 4;
	add.s64 	%rd8, %rd3, %rd15;
	ld.local.u32 	%r88, [%rd8];
	ld.local.u32 	%r89, [%rd8+-4];
	and.b32  	%r12, %r3, 31;
	setp.eq.s32	%p5, %r12, 0;
	@%p5 bra 	BB3_8;

	mov.u32 	%r53, 32;
	sub.s32 	%r54, %r53, %r12;
	shr.u32 	%r55, %r89, %r54;
	shl.b32 	%r56, %r88, %r12;
	add.s32 	%r88, %r55, %r56;
	ld.local.u32 	%r57, [%rd8+-8];
	shr.u32 	%r58, %r57, %r54;
	shl.b32 	%r59, %r89, %r12;
	add.s32 	%r89, %r58, %r59;

BB3_8:
	shr.u32 	%r60, %r89, 30;
	shl.b32 	%r61, %r88, 2;
	add.s32 	%r90, %r60, %r61;
	shl.b32 	%r18, %r89, 2;
	shr.u32 	%r62, %r90, 31;
	shr.u32 	%r63, %r88, 30;
	add.s32 	%r19, %r62, %r63;
	setp.eq.s32	%p6, %r62, 0;
	mov.u32 	%r91, %r9;
	mov.u32 	%r92, %r18;
	@%p6 bra 	BB3_10;

	not.b32 	%r64, %r90;
	neg.s32 	%r20, %r18;
	setp.eq.s32	%p7, %r18, 0;
	selp.u32	%r65, 1, 0, %p7;
	add.s32 	%r90, %r65, %r64;
	xor.b32  	%r22, %r9, -2147483648;
	mov.u32 	%r91, %r22;
	mov.u32 	%r92, %r20;

BB3_10:
	mov.u32 	%r24, %r91;
	neg.s32 	%r66, %r19;
	setp.eq.s32	%p8, %r9, 0;
	selp.b32	%r95, %r19, %r66, %p8;
	clz.b32 	%r94, %r90;
	setp.eq.s32	%p9, %r94, 0;
	shl.b32 	%r67, %r90, %r94;
	mov.u32 	%r68, 32;
	sub.s32 	%r69, %r68, %r94;
	shr.u32 	%r70, %r92, %r69;
	add.s32 	%r71, %r70, %r67;
	selp.b32	%r28, %r90, %r71, %p9;
	mov.u32 	%r72, -921707870;
	mul.hi.u32 	%r93, %r28, %r72;
	setp.lt.s32	%p10, %r93, 1;
	@%p10 bra 	BB3_12;

	mul.lo.s32 	%r73, %r28, -921707870;
	shr.u32 	%r74, %r73, 31;
	shl.b32 	%r75, %r93, 1;
	add.s32 	%r93, %r74, %r75;
	add.s32 	%r94, %r94, 1;

BB3_12:
	mov.u32 	%r76, 126;
	sub.s32 	%r77, %r76, %r94;
	shl.b32 	%r78, %r77, 23;
	add.s32 	%r79, %r93, 1;
	shr.u32 	%r80, %r79, 7;
	add.s32 	%r81, %r80, 1;
	shr.u32 	%r82, %r81, 1;
	add.s32 	%r83, %r82, %r78;
	or.b32  	%r84, %r83, %r24;
	mov.b32 	 %f44, %r84;

BB3_13:
	mul.rn.f32 	%f7, %f44, %f44;
	and.b32  	%r35, %r95, 1;
	setp.eq.s32	%p11, %r35, 0;
	@%p11 bra 	BB3_15;

	mov.f32 	%f30, 0fBAB6061A;
	mov.f32 	%f31, 0f37CCF5CE;
	fma.rn.f32 	%f45, %f31, %f7, %f30;
	bra.uni 	BB3_16;

BB3_15:
	mov.f32 	%f32, 0f3C08839E;
	mov.f32 	%f33, 0fB94CA1F9;
	fma.rn.f32 	%f45, %f33, %f7, %f32;

BB3_16:
	@%p11 bra 	BB3_18;

	mov.f32 	%f34, 0f3D2AAAA5;
	fma.rn.f32 	%f35, %f45, %f7, %f34;
	mov.f32 	%f36, 0fBF000000;
	fma.rn.f32 	%f46, %f35, %f7, %f36;
	bra.uni 	BB3_19;

BB3_18:
	mov.f32 	%f37, 0fBE2AAAA3;
	fma.rn.f32 	%f38, %f45, %f7, %f37;
	mov.f32 	%f39, 0f00000000;
	fma.rn.f32 	%f46, %f38, %f7, %f39;

BB3_19:
	fma.rn.f32 	%f47, %f46, %f44, %f44;
	@%p11 bra 	BB3_21;

	mov.f32 	%f40, 0f3F800000;
	fma.rn.f32 	%f47, %f46, %f7, %f40;

BB3_21:
	and.b32  	%r85, %r95, 2;
	setp.eq.s32	%p14, %r85, 0;
	@%p14 bra 	BB3_23;

	mov.f32 	%f41, 0f00000000;
	mov.f32 	%f42, 0fBF800000;
	fma.rn.f32 	%f47, %f47, %f42, %f41;

BB3_23:
	st.global.f32 	[%rd2], %f47;

BB3_24:
	ret;
}

	// .globl	clipPositive
.visible .entry clipPositive(
	.param .u64 clipPositive_param_0,
	.param .u64 clipPositive_param_1
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<4>;
	.reg .b32 	%r<5>;
	.reg .b64 	%rd<7>;


	ld.param.u64 	%rd2, [clipPositive_param_0];
	ld.param.u64 	%rd3, [clipPositive_param_1];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd3;
	@%p1 bra 	BB4_2;

	cvta.to.global.u64 	%rd4, %rd2;
	shl.b64 	%rd5, %rd1, 2;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f1, [%rd6];
	mov.f32 	%f2, 0f00000000;
	max.f32 	%f3, %f2, %f1;
	st.global.f32 	[%rd6], %f3;

BB4_2:
	ret;
}

	// .globl	shiftRandUniform
.visible .entry shiftRandUniform(
	.param .u64 shiftRandUniform_param_0,
	.param .u64 shiftRandUniform_param_1
)
{
	.reg .pred 	%p<3>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<6>;
	.reg .b64 	%rd<7>;


	ld.param.u64 	%rd3, [shiftRandUniform_param_0];
	ld.param.u64 	%rd4, [shiftRandUniform_param_1];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd4;
	@%p1 bra 	BB5_3;

	cvta.to.global.u64 	%rd5, %rd3;
	shl.b64 	%rd6, %rd1, 2;
	add.s64 	%rd2, %rd5, %rd6;
	ld.global.f32 	%f1, [%rd2];
	setp.neu.f32	%p2, %f1, 0f3F800000;
	@%p2 bra 	BB5_3;

	mov.u32 	%r5, 0;
	st.global.u32 	[%rd2], %r5;

BB5_3:
	ret;
}

	// .globl	uniformToBernoulli
.visible .entry uniformToBernoulli(
	.param .u64 uniformToBernoulli_param_0,
	.param .u64 uniformToBernoulli_param_1
)
{
	.reg .pred 	%p<3>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<7>;
	.reg .b64 	%rd<7>;


	ld.param.u64 	%rd3, [uniformToBernoulli_param_0];
	ld.param.u64 	%rd4, [uniformToBernoulli_param_1];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd4;
	@%p1 bra 	BB6_4;

	cvta.to.global.u64 	%rd5, %rd3;
	shl.b64 	%rd6, %rd1, 2;
	add.s64 	%rd2, %rd5, %rd6;
	ld.global.f32 	%f1, [%rd2];
	setp.gt.f32	%p2, %f1, 0f3F000000;
	@%p2 bra 	BB6_3;
	bra.uni 	BB6_2;

BB6_3:
	mov.u32 	%r6, 1065353216;
	st.global.u32 	[%rd2], %r6;
	bra.uni 	BB6_4;

BB6_2:
	mov.u32 	%r5, 0;
	st.global.u32 	[%rd2], %r5;

BB6_4:
	ret;
}

	// .globl	addRepeated
.visible .entry addRepeated(
	.param .u64 addRepeated_param_0,
	.param .u64 addRepeated_param_1,
	.param .u64 addRepeated_param_2,
	.param .u64 addRepeated_param_3
)
{
	.reg .pred 	%p<3>;
	.reg .f32 	%f<4>;
	.reg .b32 	%r<8>;
	.reg .b64 	%rd<17>;


	ld.param.u64 	%rd5, [addRepeated_param_0];
	ld.param.u64 	%rd6, [addRepeated_param_1];
	ld.param.u64 	%rd8, [addRepeated_param_2];
	ld.param.u64 	%rd7, [addRepeated_param_3];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd8;
	@%p1 bra 	BB7_5;

	and.b64  	%rd9, %rd7, -4294967296;
	setp.eq.s64	%p2, %rd9, 0;
	@%p2 bra 	BB7_3;

	rem.u64 	%rd16, %rd1, %rd7;
	bra.uni 	BB7_4;

BB7_3:
	cvt.u32.u64	%r5, %rd7;
	cvt.u32.u64	%r6, %rd1;
	rem.u32 	%r7, %r6, %r5;
	cvt.u64.u32	%rd16, %r7;

BB7_4:
	cvta.to.global.u64 	%rd10, %rd6;
	cvta.to.global.u64 	%rd11, %rd5;
	shl.b64 	%rd12, %rd1, 2;
	add.s64 	%rd13, %rd11, %rd12;
	ld.global.f32 	%f1, [%rd13];
	shl.b64 	%rd14, %rd16, 2;
	add.s64 	%rd15, %rd10, %rd14;
	ld.global.f32 	%f2, [%rd15];
	add.f32 	%f3, %f2, %f1;
	st.global.f32 	[%rd13], %f3;

BB7_5:
	ret;
}

	// .globl	addRepeatedPow2
.visible .entry addRepeatedPow2(
	.param .u64 addRepeatedPow2_param_0,
	.param .u64 addRepeatedPow2_param_1,
	.param .u64 addRepeatedPow2_param_2,
	.param .u64 addRepeatedPow2_param_3
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<4>;
	.reg .b32 	%r<5>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd2, [addRepeatedPow2_param_0];
	ld.param.u64 	%rd3, [addRepeatedPow2_param_1];
	ld.param.u64 	%rd5, [addRepeatedPow2_param_2];
	ld.param.u64 	%rd4, [addRepeatedPow2_param_3];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd5;
	@%p1 bra 	BB8_2;

	cvta.to.global.u64 	%rd6, %rd3;
	and.b64  	%rd7, %rd1, %rd4;
	shl.b64 	%rd8, %rd7, 2;
	add.s64 	%rd9, %rd6, %rd8;
	cvta.to.global.u64 	%rd10, %rd2;
	shl.b64 	%rd11, %rd1, 2;
	add.s64 	%rd12, %rd10, %rd11;
	ld.global.f32 	%f1, [%rd12];
	ld.global.f32 	%f2, [%rd9];
	add.f32 	%f3, %f2, %f1;
	st.global.f32 	[%rd12], %f3;

BB8_2:
	ret;
}

	// .globl	addScaler
.visible .entry addScaler(
	.param .f32 addScaler_param_0,
	.param .u64 addScaler_param_1,
	.param .u64 addScaler_param_2
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<4>;
	.reg .b32 	%r<5>;
	.reg .b64 	%rd<7>;


	ld.param.f32 	%f1, [addScaler_param_0];
	ld.param.u64 	%rd2, [addScaler_param_1];
	ld.param.u64 	%rd3, [addScaler_param_2];
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	cvt.u64.u32	%rd1, %r4;
	setp.ge.u64	%p1, %rd1, %rd3;
	@%p1 bra 	BB9_2;

	cvta.to.global.u64 	%rd4, %rd2;
	shl.b64 	%rd5, %rd1, 2;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f2, [%rd6];
	add.f32 	%f3, %f2, %f1;
	st.global.f32 	[%rd6], %f3;

BB9_2:
	ret;
}


`
