// 将当前活动图层移动到画布右边距 30px 的位置
var doc = app.activeDocument;
var layer = doc.activeLayer;

// 图层边界（单位：像素）
var bounds = layer.bounds; 
var layerWidth = bounds[2].as("px") - bounds[0].as("px");

// 目标X坐标 = 画布宽度 - 30 - 图层宽度
var targetX = doc.width.as("px") - 30 - layerWidth;

// 当前X = 图层左边界
var currentX = bounds[0].as("px");

// 偏移量
var deltaX = targetX - currentX;

// 移动图层
layer.translate(deltaX, 0);