/**
 * 批量导入：依次打开文件夹图片 → 合并可见 → 复制并粘贴到【当前激活的目标文档】
 * 每导入一次：目标画布向右相对扩展 33px，并把新图层贴到右边距 30px
 * 用法：先打开/新建并激活【目标文档】，再运行脚本并选择源文件夹
 */
(function () {
    if (app.documents.length < 1) { alert("请先打开一个【目标文档】并激活它。"); return; }
    var dst = app.activeDocument;

    // 选择源文件夹
    var folder = Folder.selectDialog("选择要批处理的源文件夹");
    if (!folder) return;

    // 过滤支持的文件
    var files = folder.getFiles(function (f) {
        return f instanceof File && /\.(psd|psb|png|jpe?g|tif?f)$/i.test(f.name);
    });
    if (!files || files.length === 0) { alert("该文件夹没有可处理的图片文件。"); return; }

    // 可选：按文件名排序
    files.sort(function (a, b) { return a.name.toLowerCase() > b.name.toLowerCase() ? 1 : -1; });

    // 保护性设置
    var origUnits = app.preferences.rulerUnits; app.preferences.rulerUnits = Units.PIXELS;
    var origDialogs = app.displayDialogs; app.displayDialogs = DialogModes.NO;

    try {
        for (var i = 0; i < files.length; i++) {
            // 1) 打开源文件
            var src = app.open(files[i]);
            var srcBaseName = src.name.replace(/\.[^\.]+$/, "");

            // 2) 合并可见
            try { src.mergeVisibleLayers(); } catch (e) {}

            // 3) 复制合并像素
            src.selection.selectAll();
            src.selection.copy(true);

            // 4) 关闭源文件（不保存）
            src.close(SaveOptions.DONOTSAVECHANGES);

            // 5) 激活目标文档（确保它是最前面的文档）
            app.activeDocument = dst;

            // 6) 粘贴并重命名图层
            var pasted = dst.paste();
            pasted.name = srcBaseName;

            // 7) 目标画布向右相对扩展 33px（锚点左中）
            var newWidthPx = dst.width.as("px") + 35;
            dst.resizeCanvas(new UnitValue(newWidthPx, "px"), dst.height, AnchorPosition.MIDDLELEFT);


            // 8) 把新图层贴到右边距 30px
            var b = pasted.bounds;
            var currentX = b[0].as("px");
            var targetX = dst.width.as("px") - 35;
            pasted.translate(targetX - currentX, 0);
        }

        alert("完成：共处理 " + files.length + " 个文件。");
    } catch (err) {
        alert("发生错误：\n" + err);
    } finally {
        app.displayDialogs = origDialogs;
        app.preferences.rulerUnits = origUnits;
    }
})();