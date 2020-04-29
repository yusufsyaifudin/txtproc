package similarity

import (
	"context"
	"testing"

	"github.com/opentracing/opentracing-go"
)

func TestExactMatcher_Compare1(t *testing.T) {
	m := ExactMatcher()
	score, _ := m.Compare(context.Background(), "", "x")
	if score != 0 {
		t.Error("ExactMatcher must return 0 on different string")
		t.Fail()
	}
}

func TestExactMatcher_Compare2(t *testing.T) {
	m := ExactMatcher()
	score, _ := m.Compare(context.Background(), "x", "x")
	if score != 1 {
		t.Error("ExactMatcher must return 1 on similar string")
		t.Fail()
	}
}

func TestExactMatcher(t *testing.T) {
	if ExactMatcher() == nil {
		t.Error("ExactMatcher should return not nil")
		t.Fail()
	}
}

func compareRune(ctx context.Context, str1, str2 string) bool {
	span, ctx := opentracing.StartSpanFromContext(ctx, "compareRune")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	if len(str1) != len(str2) {
		return false
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			continue
		}

		return false
	}

	return true
}

func compareEqual(ctx context.Context, str1, str2 string) bool {
	span, ctx := opentracing.StartSpanFromContext(ctx, "compareEqual")
	defer func() {
		ctx.Done()
		span.Finish()
	}()
	if str1 == str2 {
		return true
	}

	return false
}

func Benchmark_compareRune_1Word(b *testing.B) {
	ctx := context.Background()
	text := `word`

	for n := 0; n < b.N; n++ {
		if !compareRune(ctx, text, text) {
			b.Error("ExactMatcher must return 1 on similar string")
			b.Fail()
		}
	}
}

func Benchmark_compareEqual_1Word(b *testing.B) {
	text := `word`
	ctx := context.Background()

	for n := 0; n < b.N; n++ {
		if !compareEqual(ctx, text, text) {
			b.Error("ExactMatcher must return 1 on similar string")
			b.Fail()
		}
	}
}

func Benchmark_compareRune_100Word(b *testing.B) {
	ctx := context.Background()
	text := `accumsan tortor posuere ac ut consequat semper viverra nam libero justo laoreet sit amet cursus sit amet dictum sit amet justo donec enim diam vulputate ut pharetra sit amet aliquam id diam maecenas ultricies mi eget mauris pharetra et ultrices neque ornare aenean euismod elementum nisi quis eleifend quam adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus urna neque viverra justo nec ultrices dui sapien eget mi proin sed libero enim sed faucibus turpis in eu mi bibendum neque egestas congue quisque egestas diam in arcu cursus euismod quis viverra nibh cras pulvinar mattis nunc sed blandit libero volutpat sed`

	for n := 0; n < b.N; n++ {
		if !compareRune(ctx, text, text) {
			b.Error("ExactMatcher must return 1 on similar string")
			b.Fail()
		}
	}
}

func Benchmark_compareEqual_100Word(b *testing.B) {
	text := `accumsan tortor posuere ac ut consequat semper viverra nam libero justo laoreet sit amet cursus sit amet dictum sit amet justo donec enim diam vulputate ut pharetra sit amet aliquam id diam maecenas ultricies mi eget mauris pharetra et ultrices neque ornare aenean euismod elementum nisi quis eleifend quam adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus urna neque viverra justo nec ultrices dui sapien eget mi proin sed libero enim sed faucibus turpis in eu mi bibendum neque egestas congue quisque egestas diam in arcu cursus euismod quis viverra nibh cras pulvinar mattis nunc sed blandit libero volutpat sed`
	ctx := context.Background()

	for n := 0; n < b.N; n++ {
		if !compareEqual(ctx, text, text) {
			b.Error("ExactMatcher must return 1 on similar string")
			b.Fail()
		}
	}
}

func Benchmark_compareRune_200Word(b *testing.B) {
	ctx := context.Background()
	text := `fringilla ut morbi tincidunt augue interdum velit euismod in pellentesque massa placerat duis ultricies lacus sed turpis tincidunt id aliquet risus feugiat in ante metus dictum at tempor commodo ullamcorper a lacus vestibulum sed arcu non odio euismod lacinia at quis risus sed vulputate odio ut enim blandit volutpat maecenas volutpat blandit aliquam etiam erat velit scelerisque in dictum non consectetur a erat nam at lectus urna duis convallis convallis tellus id interdum velit laoreet id donec ultrices tincidunt arcu non sodales neque sodales ut etiam sit amet nisl purus in mollis nunc sed id semper risus in hendrerit gravida rutrum quisque non tellus orci ac auctor augue mauris augue neque gravida in fermentum et sollicitudin ac orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor aliquam nulla facilisi cras fermentum odio eu feugiat pretium nibh ipsum consequat nisl vel pretium lectus quam id leo in vitae turpis massa sed elementum tempus egestas sed sed risus pretium quam vulputate dignissim suspendisse in est ante in nibh mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing commodo elit at imperdiet dui accumsan sit amet nulla facilisi morbi tempus iaculis urna id volutpat lacus laoreet non curabitur gravida arcu ac`

	for n := 0; n < b.N; n++ {
		if !compareRune(ctx, text, text) {
			b.Error("ExactMatcher must return 1 on similar string")
			b.Fail()
		}
	}
}

func Benchmark_compareEqual_200Word(b *testing.B) {
	text := `fringilla ut morbi tincidunt augue interdum velit euismod in pellentesque massa placerat duis ultricies lacus sed turpis tincidunt id aliquet risus feugiat in ante metus dictum at tempor commodo ullamcorper a lacus vestibulum sed arcu non odio euismod lacinia at quis risus sed vulputate odio ut enim blandit volutpat maecenas volutpat blandit aliquam etiam erat velit scelerisque in dictum non consectetur a erat nam at lectus urna duis convallis convallis tellus id interdum velit laoreet id donec ultrices tincidunt arcu non sodales neque sodales ut etiam sit amet nisl purus in mollis nunc sed id semper risus in hendrerit gravida rutrum quisque non tellus orci ac auctor augue mauris augue neque gravida in fermentum et sollicitudin ac orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor aliquam nulla facilisi cras fermentum odio eu feugiat pretium nibh ipsum consequat nisl vel pretium lectus quam id leo in vitae turpis massa sed elementum tempus egestas sed sed risus pretium quam vulputate dignissim suspendisse in est ante in nibh mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing commodo elit at imperdiet dui accumsan sit amet nulla facilisi morbi tempus iaculis urna id volutpat lacus laoreet non curabitur gravida arcu ac`
	ctx := context.Background()

	for n := 0; n < b.N; n++ {
		if !compareEqual(ctx, text, text) {
			b.Error("ExactMatcher must return 1 on similar string")
			b.Fail()
		}
	}
}
