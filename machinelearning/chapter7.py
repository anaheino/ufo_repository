#!/home/anah/miniconda3/bin/python
from fastai.vision.all import *
import torch

path = untar_data(URLs.IMAGENETTE)
# We're loading the data into dataloaders, using previous chapters presizing-technique.
# This allows us to convert all images to same size (bs=20 is all my localhost can currently handle)
dblock = DataBlock(blocks=(ImageBlock(), CategoryBlock()),
                   get_items=get_image_files,
                   get_y=parent_label,
                   item_tfms=Resize(460),
                   batch_tfms=aug_transforms(size=224, min_scale=0.75))
dls = dblock.dataloaders(path, bs=80)

# Implementing a baseline training run:
model = xresnet50(n_out=dls.c)
learn = Learner(dls, model, loss_func=CrossEntropyLossFlat(), metrics=accuracy)
torch.cuda.empty_cache()

# Again, commented cause rough on the gpu and takes ages to run
print("are cuda's available?")
print(torch.cuda.is_available())
print(torch.cuda.get_device_name(0))
print("Let's see")
learn.fit_one_cycle(5, 3e-3)
print("ran")
# Important trick: NORMALIZATION

# When data is normalized, it has a mean of 0 and standard deviation of 1.
# Most computer vision libraries use values between 0 and 255 for pixels (256, gee, I wonder why!)
# Or between 0 and 1.

# Grabbing batch of the data, and then look at the data: averaging over all axes except for channel axis, which is 1.
x, y = dls.one_batch()
print(x.mean(dim=[0, 2, 3]), x.std(dim=[0, 2, 3]))

# Here, we actually use fastai-library's Normalize functions. Passing this transforms the mean and sd.
# If you don't pass any statistics to the transform, fastai automatically calculates them from a single batch.
def get_dls(bs, size):
    dblock = DataBlock(blocks=(ImageBlock, CategoryBlock),
                   get_items=get_image_files,
                   get_y=parent_label,
                   item_tfms=Resize(460),
                   batch_tfms=[*aug_transforms(size=size, min_scale=0.75),
                               Normalize.from_stats(*imagenet_stats)])
    return dblock.dataloaders(path, bs=bs)

# Here we load this normalized data:
dls = get_dls(64, 224)
# Then let's check the on batch:
x, y = dls.one_batch()
print(x.mean(dim=[0, 2, 3]), x.std(dim=[0, 2, 3]))
# Now we rerun this learner to find out if anything changed:

model = xresnet50(n_out=dls.c)
learn = Learner(dls, model, loss_func=CrossEntropyLossFlat(), metrics=accuracy)
# Heavy on gpu
# learn.fit_one_cycle(5, 3e-3)

# This is very important especially when working with pre-trained models. Pre-trained model
# Might have been trained with 0-255 and doesn't work properly with 0-1.
# This means when distributing a model, also distribute the correct normalization.
# Before we used vision_learner pretrained model, thus fastai automatically added the proper normalization.

# PROTIP: Start training with smaller images! Later move on to larger images.
# This means the training is completeted much faster.
# We have to use fine_tune method though: because we are training model to do something different:
# Recognize larger images. Similar to transfer learning.

# Here we are smart: let's make 'em smaller, thus the model trains faster.
dls = get_dls(64, 128)
learn = Learner(dls, xresnet50(n_out=dls.c), loss_func=CrossEntropyLossFlat(),
                metrics=accuracy)
# learn.fit_one_cycle(4, 3e-3)
# After running this, we replace dataloaders inside the learner and fine-tune:
learn.dls = get_dls(20, 224)
# learn.fine_tune(5, 1e-3)
# Note: progressive resizing might hurt when using transfer-learning. This is because if it was
# originally trained on similar images, smaller images may harm the initial weights.

# Test time augmentation: Taking multiple random crops of each image, using data augmentation, and then
# taking average or maximum of the predictions for each of the crop of the same image.
# By default, fastai uses unaugmented center crop of the image plus 4 randomly augmented images.

# We can pass any dataloader to fastai's tta, by default it will use validation set.
preds, targs = learn.tta()
print(accuracy(preds, targs).item())
# This boosts accuracy, yet makes inference slower.

#Mixup
# Mixup shines when you don't have a lot of data and don't have a pretrained model that was trained with similar
# Dataset. Mixup does the following for each image:
# 1. Select another random image from dataset
# 2. Pick a random weight
# 3. Take a weighted average (using 2's value) of the selected image with your image: this is the independent var.
# 4. Take a weighted average (same weight) of this images labels with your images labels: dependent variable

# Pseudocode (t weight of our weighted average):
# image2, target2 = dataset[randint(0,len(dataset)]
# t = random_float(0.5,1.0)
# new_image = t * image1 + (1-t) * image2
# new_target = t * target1 + (1-t) * target2

# This is mixup, as actual code:
church = PILImage.create(get_image_files_sorted(path/'train'/'n03028079')[0])
gas = PILImage.create(get_image_files_sorted(path/'train'/'n03425413')[0])
church = church.resize((256, 256))
gas = gas.resize((256, 256))
tchurch = tensor(church).float() / 255.
tgas = tensor(gas).float() / 255.

_, axs = plt.subplots(1, 3, figsize=(12, 4))
# show_image(tchurch, ax=axs[0]);
# show_image(tgas, ax=axs[1]);
# show_image((0.3*tchurch + 0.7*tgas), ax=axs[2]);
# Third image is 0.3 times the first and 0.7 times the second. So correct answer is: 30% first, 70% second.
# Like so: [0, 0, 0.3, 0, 0, 0, 0, 0.7, 0, 0]
# Mixup requires a shitload of more training, and by rule of fist not to be used for under 80 epochs.


# Label smoothing:
# Replace 1̈́'s with something a little less than 1, and 0's with something a little more than 0. Then train.
# Helps to deal with faulty datasets.
# in this case, as an example:
print([0.01, 0.01, 0.01, 0.91, 0.01, 0.01, 0.01, 0.01, 0.01, 0.01])
# In practice, one-hot encoding labels is not something we do, but it's a good visualization.

# Here we use this technique:
model = xresnet50(n_out=dls.c)
learn = Learner(dls, model, loss_func=LabelSmoothingCrossEntropy(),
                metrics=accuracy)
# learn.fit_one_cycle(5, 3e-3)

# But similar to mixup, this requires a lot of epochs.