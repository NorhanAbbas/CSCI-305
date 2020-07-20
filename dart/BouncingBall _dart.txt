import 'dart:html';
import 'dart:async';
import 'dart:math' as math;

//import 'package:universal_html/html.dart' as html;
//import 'dart:math' as Math;

CanvasElement canvas;
CanvasRenderingContext2D ctx;

Ball ball;

class Ball {
  double x = 50;
  double y = 50;
  double r = 15;
  double vx = 0;
  double vy = 0;
  Ball(this.x, this.y);
}

void clear() {
  ctx.save();
  ctx.fillStyle = "rgba(200, 0, 200, .3)";
  ctx.fillRect(0, 0, canvas.width, canvas.height);
  ctx.restore();
}

void update(ball) {
// Update ball (with Physics! =)

// 1 - apply velocity to position (vx -> x)
  ball.x += ball.vx;
  ball.y += ball.vy;

// 2 - apply drag/friction to velocity
  ball.vx *= .99;
  ball.vy *= .99;
// 3 - apply gravity to velocity
  ball.vy += .25;
  ball.vx += .25;
// 4 - check for collisions
  if (ball.y + ball.r > canvas.height) {
    ball.y = canvas.height - ball.r;
//ball.vy = -Math.abs(ball.vy);
    ball.vy = -ball.vy.abs();
  }
  if (ball.x + ball.r > canvas.width) {
    ball.x = canvas.width - ball.r;
//ball.vx = -Math.abs(ball.vx);
    ball.vx = -ball.vx.abs();
  }
// Draw ball
  clear();
  ctx.save();
  
  ctx.translate(ball.x, ball.y);
  ctx.fillStyle = "#ffff00";
  ctx.beginPath();
  ctx.arc(0, 0, ball.r, 0, 44 / 7, true);
  ctx.closePath();
  ctx.fill();
  ctx.restore();
}

void jump(MouseEvent event) {
  ball.x = event.offset.x;
  ball.y = event.offset.y;
  update(ball);
}

void main() {
  //var header = querySelector('#header');
  //header.text = "Hello, World!";

  //document.getElementById('button').clicked == true
  canvas = querySelector('#canvas');
  ctx = canvas.getContext('2d');
  ball = new Ball(1, 1);

  //new Timer.periodic(new Duration(milliseconds: 12), update);
  //clear();
  
  const move_periodically = Duration(milliseconds: 12);
  new Timer.periodic(move_periodically, (Timer t) => update(ball));
  clear();

  canvas.onClick.listen(jump);
  clear();

  //ball;
  //for (var i = 0; i < 4; i++) {
  //print('hello $i');
  //}
}

